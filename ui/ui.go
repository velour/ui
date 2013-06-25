// Package ui provides a higher-level user-interface API using SDL2.
package ui

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"time"

	"github.com/eaburns/sdl2"
)

func init() {
	runtime.LockOSThread()
}

var (
	doChan       = make(chan func(), 1)
	windows      = make(map[sdl2.WindowID]*Window, 1)
	missedEvents []windowIDer
)

// Hijack must be called by the main go routine to hijack it for the user interface;
// it never returns. The function f is called in a new go routine as the new "main" function.
// Rate is the rate at which the user interface should poll for events.
func Hijack(f func(), rate time.Duration) {
	if err := sdl2.Init(sdl2.Everything); err != nil {
		panic(err)
	}

	go f()

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

type windowIDer interface {
	WindowID() sdl2.WindowID
}

func pollEvents() {
	for {
		missed := missedEvents
		missedEvents = nil
		for _, m := range missed {
			sendEvent(m)
		}

		ev := sdl2.PollEvent()
		if ev == nil {
			break
		}
		if e, ok := ev.(windowIDer); ok {
			sendEvent(e)
		}
	}
}

func sendEvent(e windowIDer) {
	if win, ok := windows[e.WindowID()]; ok {
		win.events <- e
	} else {
		// Must be that the window is still being created.  Let's try again later.
		missedEvents = append(missedEvents, e)
	}
}

// A Window is a single window on the user's graphical interface.
type Window struct {
	win    *sdl2.Window
	rend   *sdl2.Renderer
	events chan interface{}
	imgs   map[string]img
}

type img struct {
	tex           *sdl2.Texture
	width, height int
}

// NewWindow returns a new window.
func NewWindow(title string, w, h int) *Window {
	win := &Window{
		events: make(chan interface{}, 1),
		imgs:   make(map[string]img),
	}
	do(func() {
		var err error
		x, y := sdl2.WindowPosUndefined, sdl2.WindowPosUndefined
		win.win, err = sdl2.CreateWindow(title, x, y, w, h, sdl2.Shown|sdl2.OpenGL)
		if err != nil {
			panic(err)
		}

		win.rend, err = sdl2.CreateRenderer(win.win, -1, sdl2.Accelerated)
		if err != nil {
			panic(err)
		}
		if err := win.rend.SetDrawBlendMode(sdl2.Blend); err != nil {
			panic(err)
		}

		windows[win.win.ID()] = win
	})
	return win
}

// Destroy destroys the window.
func (win *Window) Destroy() {
	do(func() {
		win.rend.Destroy()
		win.win.Destroy()
	})
}

// FlushCache flushes any cached textures on this window.
func (win *Window) FlushCache() {
	do(func() {
		for _, img := range win.imgs {
			img.tex.Destroy()
		}
		win.imgs = make(map[string]img)
	})
}

// Events returns the event channel for the window.
func (win *Window) Events() <-chan interface{} {
	return win.events
}

// Draw calls f from the main go routine. F is passed a canvas which can draw to the
// window. The Canvas's methods can only safely be called from the main go routine.
func (win *Window) Draw(f func(win Canvas)) {
	do(func() { f(Canvas{win}) })
}

// A Canvas can draw to a window.
// The drawing operations can only be safely used within the main go routine.
type Canvas struct {
	win *Window
}

// Clear clears the canvas with the drawing color.
func (canv Canvas) Clear() {
	must(canv.win.rend.Clear())
}

// DrawLine draws a line on the canvas.
func (canv Canvas) DrawLine(x1, y1, x2, y2 int) {
	must(canv.win.rend.DrawLine(x1, y1, x2, y2))
}

// DrawLines draws a series of connected lines on the canvas.
func (canv Canvas) DrawLines(points []image.Point) {
	must(canv.win.rend.DrawLines(points))
}

// DrawPoint draws a point on the canvas.
func (canv Canvas) DrawPoint(x, y int) {
	must(canv.win.rend.DrawPoint(x, y))
}

// DrawPoints draws multiple points on the canvas.
func (canv Canvas) DrawPoints(points []image.Point) {
	must(canv.win.rend.DrawPoints(points))
}

// DrawRect draws a rectangle on the canvas.
func (canv Canvas) DrawRect(rect *image.Rectangle) {
	must(canv.win.rend.DrawRect(rect))
}

// DrawRects draws some number of rectangles on the canvas.
func (canv Canvas) DrawRects(rects []image.Rectangle) {
	must(canv.win.rend.DrawRects(rects))
}

// FillRect fills a rectangle on the canvas with the drawing color.
func (canv Canvas) FillRect(rect *image.Rectangle) {
	must(canv.win.rend.FillRect(rect))
}

// FillRects fills some number of rectangles on the canvas with the drawing color.
func (canv Canvas) FillRects(rects []image.Rectangle) {
	must(canv.win.rend.FillRects(rects))
}

// Present updates the canvas with drawing performed.
func (canv Canvas) Present() {
	canv.win.rend.Present()
}

// SetDrawColor sets the color used for drawing operations (Rect, Line and Clear).
func (canv Canvas) SetDrawColor(col color.Color) {
	must(canv.win.rend.SetDrawColor(col))
}

// DrawPNG draws the image loaded from a PNG file to the canvas.
// The image is drawn with the upper-left corner located at x, y.
func (canv Canvas) DrawPNG(path string, x, y int) {
	i, ok := canv.win.imgs[path]
	if !ok {
		png := loadPNG(path)
		tex, err := sdl2.CreateTextureFromImage(canv.win.rend, png)
		if err != nil {
			panic(err)
		}
		if err := tex.SetBlendMode(sdl2.Blend); err != nil {
			panic(err)
		}
		i = img{
			tex:    tex,
			width:  png.Bounds().Dx(),
			height: png.Bounds().Dy(),
		}
		canv.win.imgs[path] = i
	}
	dst := image.Rect(x, y, x+i.width, y+i.height)
	must(canv.win.rend.Copy(i.tex, nil, &dst))
}

func loadPNG(path string) *image.NRGBA {
	r, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}
	return img.(*image.NRGBA)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
