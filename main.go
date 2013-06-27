// +build ignore

package main

import (
	"image/color"
	"time"

	"github.com/eaburns/ui"
)

const (
	width   = 640
	height  = 480
	imgPath = "gopher.png"
)

func main() {
	ui.Start(main2, 20*time.Millisecond)
}

func main2() {
	win := ui.NewWindow("test", width, height)
	tick := time.NewTicker(20 * time.Millisecond)
	for {
		select {
		case ev := <-win.Events():
			if w, ok := ev.(*ui.WindowEvent); ok && w.Event == ui.WindowClose {
				return
			}
		case <-tick.C:
			win.Draw(func(w ui.Canvas) {
				w.SetDrawColor(color.Black)
				w.Clear()
				w.DrawPNG(imgPath, 0, 0)
			})
		}
	}
}
