// +build ignore

package main

import (
	"image/color"
	"os"
	"time"

	"github.com/eaburns/sdl2"
	"github.com/eaburns/sdl2/ui"
)

const (
	width   = 640
	height  = 480
	imgPath = "gopher.png"
)

func main() {
	ui.Hijack(main2, 20*time.Millisecond)
}

func main2() {
	win := ui.NewWindow("test", width, height)
	tick := time.NewTicker(20 * time.Millisecond)
	for {
		select {
		case ev := <-win.Events():
			if w, ok := ev.(*sdl2.WindowEvent); ok && w.Event == sdl2.WindowClose {
				os.Exit(0)
			}
		case <-tick.C:
			win.Draw(func(w ui.Canvas) {
				w.SetDrawColor(color.Black)
				w.Clear()
				w.DrawPNG(imgPath, 0, 0)
				w.Present()
			})
		}
	}
}
