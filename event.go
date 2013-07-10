// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include <SDL.h>

#cgo darwin CFLAGS: -I/Library/Frameworks/SDL2.framework/Headers
#cgo darwin LDFLAGS: -framework SDL2

#cgo linux pkg-config: sdl2

Uint32 eventType(SDL_Event *e) {
	return e->type;
}
*/
import "C"

import (
	"strconv"
	"unsafe"
)

// PollEvent polls for currently pending events.
//
// Events not supported by this binding set are discarded.
func pollEvent() interface{} {
	for {
		var ev C.SDL_Event
		if C.SDL_PollEvent(&ev) == 0 {
			return nil
		}

		switch C.eventType(&ev) {
		case C.SDL_KEYDOWN, C.SDL_KEYUP:
			return newKeyboardEvent(&ev)
		case C.SDL_WINDOWEVENT:
			return newWindowEvent(&ev)
		case C.SDL_MOUSEWHEEL:
			return newMouseWheelEvent(&ev)
		case C.SDL_MOUSEMOTION:
			return newMouseMotionEvent(&ev)
		case C.SDL_MOUSEBUTTONDOWN, C.SDL_MOUSEBUTTONUP:
			return newMouseButtonEvent(&ev)
		}
	}
}

type WindowEventKind C.Uint32

const (
	// WindowShown says that the window has been shown.
	WindowShown WindowEventKind = C.SDL_WINDOWEVENT_SHOWN

	// WindowHidden says that the window has been hidden.
	WindowHidden WindowEventKind = C.SDL_WINDOWEVENT_HIDDEN

	// WindowExposed says that the window has been exposed and should be redrawn.
	WindowExposed WindowEventKind = C.SDL_WINDOWEVENT_EXPOSED

	// WindowMoved says that the window has been moved to Data1, Data2.
	WindowMoved WindowEventKind = C.SDL_WINDOWEVENT_MOVED

	// WindowResized says that the window has been resized to Data1xData2,
	WindowResized WindowEventKind = C.SDL_WINDOWEVENT_RESIZED

	// WindowSizeChanged says that the window size has changed, either as a result
	// of an API call or through the system or user changing the window size.
	WindowSizeChanged WindowEventKind = C.SDL_WINDOWEVENT_SIZE_CHANGED

	// WindowMinimized says that the window has been minimized.
	WindowMinimized WindowEventKind = C.SDL_WINDOWEVENT_MINIMIZED

	// WindowMaximized says that the window has been maximized.
	WindowMaximized WindowEventKind = C.SDL_WINDOWEVENT_MAXIMIZED

	// WindowRestored says that the window has been restored to normal size and position.
	WindowRestored WindowEventKind = C.SDL_WINDOWEVENT_RESTORED

	// WindowEnter says that the window has gained mouse focus.
	WindowEnter WindowEventKind = C.SDL_WINDOWEVENT_ENTER

	// WindowLeave says that the window has lost mouse focus.
	WindowLeave WindowEventKind = C.SDL_WINDOWEVENT_LEAVE

	// WindowFocusGained says that the window has gained keyboard focus.
	WindowFocusGained WindowEventKind = C.SDL_WINDOWEVENT_FOCUS_GAINED

	// WindowFocusLost says that the window has lost keyboard focus.
	WindowFocusLost WindowEventKind = C.SDL_WINDOWEVENT_FOCUS_LOST

	// WindowClose says that the window manager requests that the window be closed.
	WindowClose WindowEventKind = C.SDL_WINDOWEVENT_CLOSE
)

var windowEventKindNames = map[WindowEventKind]string{
	WindowShown:       "WindowShown",
	WindowHidden:      "WindowHidden",
	WindowExposed:     "WindowExposed",
	WindowMoved:       "WindowMoved",
	WindowResized:     "WindowResized",
	WindowSizeChanged: "WindowSizeChanged",
	WindowMinimized:   "WindowMinimized",
	WindowMaximized:   "WindowMaximized",
	WindowRestored:    "WindowRestored",
	WindowEnter:       "WindowEnter",
	WindowLeave:       "WindowLeave",
	WindowFocusGained: "WindowFocusGained",
	WindowFocusLost:   "WindowFocusLost",
	WindowClose:       "WindowClose",
}

func (w WindowEventKind) String() string {
	if n, ok := windowEventKindNames[w]; ok {
		return n
	}
	return "Unknown(" + strconv.Itoa(int(w)) + ")"
}

// A WindowEvent is a structure that contains window state change event data.
type WindowEvent struct {
	winID        windowID
	Event        WindowEventKind
	Data1, Data2 int
}

func (e WindowEvent) windowID() windowID {
	return e.winID
}

func newWindowEvent(ev *C.SDL_Event) *WindowEvent {
	e := (*C.SDL_WindowEvent)(unsafe.Pointer(ev))
	return &WindowEvent{
		winID: windowID(e.windowID),
		Event: WindowEventKind(e.event),
		Data1: int(e.data1),
		Data2: int(e.data2),
	}
}

type KeyboardEvent struct {
	winID  windowID
	Down   bool
	Repeat bool
	Key    Key
}

func (e KeyboardEvent) windowID() windowID {
	return e.winID
}

func newKeyboardEvent(ev *C.SDL_Event) *KeyboardEvent {
	e := (*C.SDL_KeyboardEvent)(unsafe.Pointer(ev))
	return &KeyboardEvent{
		winID:  windowID(e.windowID),
		Down:   e.state == C.SDL_PRESSED,
		Repeat: e.repeat != 0,
		Key:    Key(e.keysym.sym),
	}
}

type MouseMotionEvent struct {
	winID            windowID
	X, Y, Xrel, Yrel int
}

func newMouseMotionEvent(ev *C.SDL_Event) *MouseMotionEvent {
	e := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(ev))
	return &MouseMotionEvent{
		winID: windowID(e.windowID),
		X:     int(e.x),
		Y:     int(e.y),
		Xrel:  int(e.xrel),
		Yrel:  int(e.yrel),
	}
}

func (e MouseMotionEvent) windowID() windowID {
	return e.winID
}

type Button C.Uint8

const (
	ButtonLeft   Button = C.SDL_BUTTON_LEFT
	ButtonMiddle Button = C.SDL_BUTTON_MIDDLE
	ButtonRight  Button = C.SDL_BUTTON_RIGHT
	ButtonX1     Button = C.SDL_BUTTON_X1
	ButtonX2     Button = C.SDL_BUTTON_X2
)

var buttonNames = map[Button]string{
	ButtonLeft:   "ButtonLeft",
	ButtonMiddle: "ButtonMiddle",
	ButtonRight:  "ButtonRight",
	ButtonX1:     "ButtonX1",
	ButtonX2:     "ButtonX2",
}

func (b Button) String() string {
	if n, ok := buttonNames[b]; ok {
		return n
	}
	return "Unknown(" + strconv.Itoa(int(b)) + ")"

}

type MouseButtonEvent struct {
	winID  windowID
	Button Button
	Down   bool
	X, Y   int
}

func (e MouseButtonEvent) windowID() windowID {
	return e.winID
}

func newMouseButtonEvent(ev *C.SDL_Event) *MouseButtonEvent {
	e := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(ev))
	return &MouseButtonEvent{
		winID:  windowID(e.windowID),
		Button: Button(e.button),
		Down:   e.state == C.SDL_PRESSED,
		X:      int(e.x),
		Y:      int(e.y),
	}
}

type MouseWheelEvent struct {
	winID windowID
	X, Y  int
}

func newMouseWheelEvent(ev *C.SDL_Event) *MouseWheelEvent {
	e := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(ev))
	return &MouseWheelEvent{
		winID: windowID(e.windowID),
		X:     int(e.x),
		Y:     int(e.y),
	}
}

func (e MouseWheelEvent) windowID() windowID {
	return e.winID
}
