package sdl2

/*
#include <SDL.h>

#cgo darwin CFLAGS: -I/Library/Frameworks/SDL2.framework/Headers
#cgo darwin LDFLAGS: -framework SDL2

#cgo linux pkg-config: sdl2
*/
import "C"

import (
	"unsafe"
)

type Window C.SDL_Window

func (win *Window) sdl() *C.SDL_Window {
	return (*C.SDL_Window)(win)
}

type WindowFlags C.Uint32

const (
	FullScreen        WindowFlags = C.SDL_WINDOW_FULLSCREEN
	FullScreenDesktop WindowFlags = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	OpenGL            WindowFlags = C.SDL_WINDOW_OPENGL
	Shown             WindowFlags = C.SDL_WINDOW_SHOWN
	Hidden            WindowFlags = C.SDL_WINDOW_HIDDEN
	Borderless        WindowFlags = C.SDL_WINDOW_BORDERLESS
	Resizable         WindowFlags = C.SDL_WINDOW_RESIZABLE
	Minimized         WindowFlags = C.SDL_WINDOW_MINIMIZED
	Maximized         WindowFlags = C.SDL_WINDOW_MAXIMIZED
	InputGrabbed      WindowFlags = C.SDL_WINDOW_INPUT_GRABBED
	InputFocus        WindowFlags = C.SDL_WINDOW_INPUT_FOCUS
	MouseFocus        WindowFlags = C.SDL_WINDOW_MOUSE_FOCUS
)

const (
	// WindowPosCentered can be used as either the x or y location of the
	// CreateWindow function to specify that the window should be centered.
	WindowPosCentered = int(C.SDL_WINDOWPOS_CENTERED)

	// WindowPosUndefined can be used as either the x or y location of the
	// CreateWindow function to specify that the window should be placed
	// anywhere.
	WindowPosUndefined = int(C.SDL_WINDOWPOS_UNDEFINED)
)

// CreateWindow creates a window with the specified position, dimensions, and flags.
func CreateWindow(title string, x, y, w, h int, flags WindowFlags) (*Window, error) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	win := C.SDL_CreateWindow(ctitle, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if win == nil {
		return nil, sdlError()
	}
	return (*Window)(win), nil
}

// Destroy destroys a window.
func (win *Window) Destroy() {
	C.SDL_DestroyWindow(win.sdl())
}

// Size returns the size of a window's client area.
func (win *Window) Size() (int, int) {
	var w, h C.int
	C.SDL_GetWindowSize(win.sdl(), &w, &h)
	return int(w), int(h)
}

type WindowID C.Uint32

// ID returns the numeric ID of a window, for logging purposes.
func (win *Window) ID() WindowID {
	return WindowID(C.SDL_GetWindowID(win.sdl()))
}

// GetWindowFromID returns a window from a stored ID.
func GetWindowFromID(id WindowID) (*Window, error) {
	win := C.SDL_GetWindowFromID(C.Uint32(id))
	if win == nil {
		return nil, sdlError()
	}
	return (*Window)(win), nil
}
