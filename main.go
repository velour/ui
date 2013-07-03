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
	lastFrame := time.Now()
	var frameTime, drawTime time.Duration
	for {
		select {
		case ev := <-win.Events():
			if w, ok := ev.(*ui.WindowEvent); ok && w.Event == ui.WindowClose {
				return
			}
		case <-tick.C:
			startDraw := time.Now()
			win.Draw(func(c ui.Canvas) {
				c.SetColor(color.White)
				c.Clear()
				c.DrawPNG(imgPath, 0, 0)

				c.SetColor(color.NRGBA{G: 128, A: 255})
				c.SetFont(font, 12)
				_, h := c.FillString("Hello, World!", 50, 50)

				c.SetColor(color.NRGBA{B: 255, A: 128})
				c.SetFont(font, 48)
				w, _ := c.FillString("Foo bar", 50, 50+h)
				c.FillString(" baz", 50+w, 50+h)

				c.SetColor(color.RGBA{B: 255, G: 128, A: 255})
				c.SetFont(font, 12)
				frameStr := frameTime.String() + " frame time"
				w, h = c.StringSize(frameStr)
				c.FillString(frameStr, width-w, height-h)
				drawStr := drawTime.String() + " draw time"
				w, _ = c.StringSize(drawStr)
				c.FillString(drawStr, width-w, height-2*h)
			})
			drawTime = time.Since(startDraw)
			frameTime = time.Since(lastFrame)
			lastFrame = time.Now()
		}
	}
}
