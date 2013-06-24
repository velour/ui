// +build ignore

package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"

	"github.com/eaburns/sdl2"
)

const (
	width   = 640
	height  = 480
	imgPath = "gopher.png"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := sdl2.Init(sdl2.Everything); err != nil {
		panic(err)
	}

	x, y := sdl2.WindowPosUndefined, sdl2.WindowPosUndefined
	win, err := sdl2.CreateWindow("test", x, y, width, height, sdl2.Shown|sdl2.OpenGL)
	if err != nil {
		panic(err)
	}

	rend, err := sdl2.CreateRenderer(win, -1, sdl2.Accelerated)
	if err != nil {
		panic(err)
	}

	img := LoadPng(imgPath)
	tex, err := sdl2.CreateTextureFromImage(rend, img)
	if err != nil {
		panic(err)
	}
	tex.SetBlendMode(sdl2.Blend)

	rend.SetDrawColor(color.Black)
	rend.Clear()
	rend.Copy(tex, nil, nil)
	rend.Present()

	select {}
}

func LoadPng(path string) *image.NRGBA {
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
