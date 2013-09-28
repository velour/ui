// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

// Package ui provides a higher-level user-interface API using SDL2.
package ui

/*
#include "ui.h"
*/
import "C"

import (
	"errors"
	"os"
	"runtime"
	"time"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

const eventChanSize = 100

var (
	doChan  = make(chan func(), 1)
	windows = make(map[windowID]*Window, 1)
)

// Start starts the user interface.  It must be called by the main go routine, and it
// never returns. The function f is called in a new go routine as the new "main"
// function.  Rate is the rate at which the user interface should poll for events.
func Start(f func(), rate time.Duration) {
	if C.SDL_Init(C.SDL_INIT_EVERYTHING) < 0 {
		panic(sdlError())
	}

	initAudio()

	go func() {
		f()
		os.Exit(0)
	}()

	tick := time.NewTicker(rate)
	for {
		select {
		case f := <-doChan:
			f()
		case <-tick.C:
			pollEvents()
		}
	}
}

func do(f func()) {
	done := make(chan struct{})
	doChan <- func() {
		f()
		done <- struct{}{}
	}
	<-done
}

type windowID C.Uint32

type windowIDer interface {
	windowID() windowID
}

func pollEvents() {
	for {
		if ev := pollEvent(); ev == nil {
			break
		} else if e, ok := ev.(windowIDer); ok {
			win, ok := windows[e.windowID()]
			if !ok {
				return
			}
			select {
			case win.events <- e:
			default: // too many events queued, junk it.
			}
		}
	}
}

// A Window is a single window on the user's graphical interface.
type Window struct {
	win    *C.SDL_Window
	rend   *C.SDL_Renderer
	id     windowID
	events chan interface{}
	imgs   map[string]texture
}

type texture struct {
	tex           *C.SDL_Texture
	width, height int
}

// NewWindow returns a new window.
func NewWindow(title string, w, h int) *Window {
	win := &Window{
		events: make(chan interface{}, eventChanSize),
		imgs:   make(map[string]texture),
	}
	do(func() {
		ctitle := C.CString(title)
		defer C.free(unsafe.Pointer(ctitle))
		x, y := C.SDL_WINDOWPOS_UNDEFINED, C.SDL_WINDOWPOS_UNDEFINED
		flags := C.SDL_WINDOW_SHOWN | C.SDL_WINDOW_OPENGL

		win.win = C.SDL_CreateWindow(ctitle, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
		if win.win == nil {
			panic(sdlError())
		}

		win.rend = C.SDL_CreateRenderer(win.win, 0, C.SDL_RENDERER_ACCELERATED)
		if win.rend == nil {
			panic(sdlError())
		}
		if C.SDL_SetRenderDrawBlendMode(win.rend, C.SDL_BLENDMODE_BLEND) < 0 {
			panic(sdlError())
		}

		win.id = windowID(C.SDL_GetWindowID(win.win))
		windows[win.id] = win
	})
	return win
}

// Destroy destroys the window.
func (win *Window) Destroy() {
	do(func() {
		C.SDL_DestroyRenderer(win.rend)
		C.SDL_DestroyWindow(win.win)
		delete(windows, win.id)
	})
}

// FlushCache flushes any cached textures on this window.
func (win *Window) FlushCache() {
	do(func() {
		for _, img := range win.imgs {
			C.SDL_DestroyTexture(img.tex)
		}
		win.imgs = make(map[string]texture)
	})
}

// Events returns the event channel for the window.
func (win *Window) Events() <-chan interface{} {
	return win.events
}

// Draw calls f from the main go routine. F is passed a canvas which can draw to the
// window. The Canvas's methods can only safely be called from the main go routine.
func (win *Window) Draw(f func(win Canvas)) {
	do(func() {
		f(Canvas{win: win})
		C.SDL_RenderPresent(win.rend)
	})
}

func sdlError() error {
	return errors.New(C.GoString(C.SDL_GetError()))
}
