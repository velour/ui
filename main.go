// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

// +build ignore

package main

import (
	"image/color"
	"time"

	"github.com/velour/ui"
)

const (
	width    = 640
	height   = 480
	imgPath  = "resrc/gopher.png"
	fontPath = "resrc/prstartk.ttf"
	pewPath  = "resrc/pew.wav"
	owPath   = "resrc/ow1.wav"
)

func main() {
	ui.Start(main2, 20*time.Millisecond)
}

func main2() {
	ow := ui.PlayWAV(owPath, true)
	stopOw := time.NewTicker(5 * time.Second)
	pewTick := time.NewTicker(500 * time.Millisecond)

	win := ui.NewWindow("test", width, height)
	tick := time.NewTicker(20 * time.Millisecond)
	lastFrame := time.Now()
	var frameDur, drawDur time.Duration
	for {
		select {
		case ev := <-win.Events():
			if w, ok := ev.(*ui.WindowEvent); ok && w.Event == ui.WindowClose {
				return
			}

		case <-stopOw.C:
			ow.Stop()

		case <-pewTick.C:
			ui.PlayWAV(pewPath, false)

		case <-tick.C:
			startDraw := time.Now()
			win.Draw(func(c ui.Canvas) {
				c.SetColor(color.White)
				c.Clear()
				c.DrawPNG(imgPath, 0, 0)

				c.SetColor(color.NRGBA{G: 128, A: 255})
				c.SetFont(fontPath, 12)
				_, h := c.FillString("Hello, World!", 50, 50)

				c.SetColor(color.NRGBA{B: 255, A: 128})
				c.SetFont(fontPath, 48)
				w, _ := c.FillString("Foo bar", 50, 50+h)
				c.FillString(" baz", 50+w, 50+h)

				c.SetColor(color.RGBA{B: 255, G: 128, A: 255})
				c.SetFont(fontPath, 12)
				frameStr := frameDur.String() + " frame time"
				w, h = c.StringSize(frameStr)
				c.FillString(frameStr, width-w, height-h)
				drawStr := drawDur.String() + " draw time"
				w, _ = c.StringSize(drawStr)
				c.FillString(drawStr, width-w, height-2*h)
			})
			drawDur = time.Since(startDraw)
			frameDur = time.Since(lastFrame)
			lastFrame = time.Now()
		}
	}
}
