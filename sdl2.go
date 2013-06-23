package sdl2

/*
#include <SDL.h>

#cgo darwin CFLAGS: -I/Library/Frameworks/SDL2.framework/Headers
#cgo darwin LDFLAGS: -framework SDL2

#cgo linux pkg-config: sdl2
*/
import "C"

import (
	"errors"
)

type InitFlags C.Uint32

const (
	// Audio initializes the audio subsystem.
	Audio InitFlags = C.SDL_INIT_AUDIO

	// Video initializes the video subsystem.
	Video InitFlags = C.SDL_INIT_VIDEO

	// Everything initializes all of the above subsystems.
	Everything InitFlags = C.SDL_INIT_EVERYTHING

	// NoParachute sets SDL not to catch fatal signals.
	NoParachute InitFlags = C.SDL_INIT_NOPARACHUTE
)

// Init initializes the SDL library. This must be called before using any other SDL function.
func Init(flags InitFlags) error {
	if C.SDL_Init(C.Uint32(flags)) < 0 {
		return sdlError()
	}
	return nil
}

// Quit cleans up all initialized subsystems. You should call it upon all exit conditions.
func Quit() {
	C.SDL_Quit()
}

func sdlError() error {
	return errors.New(C.GoString(C.SDL_GetError()))
}
