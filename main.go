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
	font    = "prstartk.ttf"
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
			win.Draw(func(c ui.Canvas) {
				c.SetColor(color.White)
				c.Clear()
				c.DrawPNG(imgPath, 0, 0)

				c.SetColor(color.NRGBA{G: 128, A: 255})
				c.SetFont(font, 12)
				_, h := c.DrawString("Hello, World!", 50, 50)

				c.SetColor(color.NRGBA{B: 255, A: 128})
				c.SetFont(font, 48)
				w, _ := c.DrawString("Foo bar", 50, 50+h)
				c.DrawString(" baz", 50+w, 50+h)
			})
		}
	}
}
