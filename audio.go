// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include <SDL.h>

#cgo darwin CFLAGS: -I/Library/Frameworks/SDL2.framework/Headers
#cgo darwin LDFLAGS: -framework SDL2

#cgo linux pkg-config: sdl2

extern void audioCallback(void *, Uint8 *, int);

void setCallback(SDL_AudioSpec *s) {
	s->callback = audioCallback;
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

var (
	// OpenedSpec is the audio spec describing the format currently opened for playing.
	openedSpec C.SDL_AudioSpec

	// Sounds caches the data for all played sounds.
	sounds = map[string]*audioData{}

	// Playing is a slice of all currently-playing sounds.
	playing []*Sound
)

func initAudio() {
	var want C.SDL_AudioSpec
	want.freq = 44100
	want.format = C.AUDIO_S16LSB
	want.samples = 8096
	C.setCallback(&want)
	C.SDL_OpenAudio(&want, &openedSpec)
	C.SDL_PauseAudio(0)
}

// PlayWAV returns a Sound from a WAV file.
func PlayWAV(path string, repeat bool) *Sound {
	data, ok := sounds[path]
	if !ok {
		var err error
		if data, err = loadWAV(path); err != nil {
			panic(err)
		}
		sounds[path] = data
	}
	C.SDL_LockAudio()
	defer C.SDL_UnlockAudio()
	s := &Sound{audioData: data, repeat: repeat}
	playing = append(playing, s)
	return s
}

// A Sound is a stream of currently playing audio.
type Sound struct {
	*audioData
	pos    uintptr
	repeat bool
}

// Stop stops the sound from playing.
func (s *Sound) Stop() {
	C.SDL_LockAudio()
	defer C.SDL_UnlockAudio()
	s.repeat = false
	s.pos = 0
}

func (s *Sound) done() bool {
	return (!s.repeat && s.pos >= s.len)
}

func (s *Sound) mix(stream *C.Uint8, sz C.int) {
	vol := C.int(C.SDL_MIX_MAXVOLUME)

	left := uintptr(sz)
	dst := uintptr(unsafe.Pointer(stream))
	for left > 0 {
		data := uintptr(s.data) + uintptr(s.pos)

		n := uintptr(left)
		if s.len-s.pos < n {
			n = s.len - s.pos
		}

		C.SDL_MixAudio((*C.Uint8)(unsafe.Pointer(dst)), (*C.Uint8)(unsafe.Pointer(data)), C.Uint32(n), vol)
		left -= n
		s.pos += n
		dst += uintptr(n)

		if !s.repeat {
			break
		} else if s.pos >= s.len {
			s.pos = 0
		}
	}
}

// AudioData is sound data converted to the format of the opened audio spec.
type audioData struct {
	path string
	data unsafe.Pointer
	len  uintptr
}

// Rb is a string of C.fopen flags (read and binary) for C.SDL_RWFromFile.
var rb = C.CString("rb")

func loadWAV(path string) (*audioData, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var data *C.Uint8
	var len C.Uint32
	var s C.SDL_AudioSpec // Apparently just a buffer?  SDL just zeroes it, fills it, and returns it.
	spec := C.SDL_LoadWAV_RW(C.SDL_RWFromFile(cpath, rb), 1, &s, &data, &len)
	if spec == nil {
		return nil, sdlError()
	}

	var err error
	data, len, err = convert(spec, data, len)
	if err != nil {
		return nil, err
	}

	return &audioData{
		path: path,
		data: unsafe.Pointer(data),
		len:  uintptr(len),
	}, nil
}

// Convert converts audio to the opened audio spec, returning the data and it's length.
func convert(s *C.SDL_AudioSpec, data *C.Uint8, len C.Uint32) (*C.Uint8, C.Uint32, error) {
	var cvt C.SDL_AudioCVT
	o := openedSpec
	switch C.SDL_BuildAudioCVT(&cvt, s.format, s.channels, s.freq, o.format, o.channels, o.freq) {
	case -1:
		return nil, 0, errors.New("Cannot convert audio")
	case 0:
		return data, len, nil
	}

	buf := C.malloc(C.size_t(len) * C.size_t(cvt.len_mult))
	cvt.buf = (*C.Uint8)(buf)
	cvt.len = C.int(len)
	C.memcpy(buf, unsafe.Pointer(data), C.size_t(len))
	C.free(unsafe.Pointer(data))

	if C.SDL_ConvertAudio(&cvt) < 0 {
		return nil, 0, sdlError()
	}
	return cvt.buf, C.Uint32(cvt.len), nil
}
