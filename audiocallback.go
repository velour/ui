// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include <SDL/SDL.h>
#include <string.h>

extern void *offset(void*, int);
*/
import "C"

import (
	"unsafe"
)

//export audioCallback
func audioCallback(_ unsafe.Pointer, dst *C.Uint8, sz C.int) {
	C.SDL_memset(unsafe.Pointer(dst), C.int(openedSpec.silence), C.size_t(sz))

	for i := 0; i < len(playing); i++ {
		if playing[i].done() {
			playing[i] = playing[len(playing)-1]
			playing[len(playing)-1] = nil
			playing = playing[:len(playing)-1]
			i--
			continue
		}
		playing[i].mix(dst, sz)
	}
}
