// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include <SDL.h>

#cgo darwin CFLAGS: -I/Library/Frameworks/SDL2.framework/Headers
#cgo darwin LDFLAGS: -framework SDL2

#cgo linux pkg-config: sdl2
*/
import "C"

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"unsafe"
)

// A Canvas can draw to a window.
// The drawing operations can only be safely used within the main go routine.
type Canvas struct {
	win  *Window
	font font
}

// Clear clears the canvas with the drawing color.
func (c Canvas) Clear() {
	if C.SDL_RenderClear(c.win.rend) < 0 {
		panic(sdlError())
	}
}

// SetColor sets the color used for drawing operations (DrawPoints, DrawLines, DrawRects, FillRects, and Clear).
func (c Canvas) SetColor(col color.Color) {
	r, g, b, a := col.RGBA()
	f := 255.0 / 0xFFFF
	r8 := C.Uint8(float64(r) * f)
	g8 := C.Uint8(float64(g) * f)
	b8 := C.Uint8(float64(b) * f)
	a8 := C.Uint8(float64(a) * f)
	if C.SDL_SetRenderDrawColor(c.win.rend, r8, g8, b8, a8) < 0 {
		panic(sdlError())
	}
}

// Color returns the current drawing color.
func (c Canvas) color() color.Color {
	var r, g, b, a C.Uint8
	if C.SDL_GetRenderDrawColor(c.win.rend, &r, &g, &b, &a) < 0 {
		panic(sdlError())
	}
	return color.NRGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

// DrawPoints draws multiple points on the canvas.
func (c Canvas) DrawPoints(points ...image.Point) {
	if C.SDL_RenderDrawPoints(c.win.rend, sdlPoints(points), C.int(len(points))) < 0 {
		panic(sdlError())
	}
}

// DrawLines draws a series of connected lines on the canvas.
func (c Canvas) DrawLines(points ...image.Point) {
	if C.SDL_RenderDrawLines(c.win.rend, sdlPoints(points), C.int(len(points))) < 0 {
		panic(sdlError())
	}
}

func sdlPoints(points []image.Point) *C.SDL_Point {
	pts := make([]C.SDL_Point, len(points))
	for i, p := range points {
		pts[i].x = C.int(p.X)
		pts[i].y = C.int(p.Y)
	}
	return &pts[0]
}

// DrawRects draws some number of rectangles on the canvas.
func (c Canvas) DrawRects(rects ...image.Rectangle) {
	if C.SDL_RenderDrawRects(c.win.rend, sdlRects(rects), C.int(len(rects))) < 0 {
		panic(sdlError())
	}
}

// FillRects fills some number of rectangles on the canvas with the drawing color.
func (c Canvas) FillRects(rects ...image.Rectangle) {
	if C.SDL_RenderFillRects(c.win.rend, sdlRects(rects), C.int(len(rects))) < 0 {
		panic(sdlError())
	}
}

func sdlRects(rects []image.Rectangle) *C.SDL_Rect {
	rs := make([]C.SDL_Rect, len(rects))
	for i := range rects {
		rs[i] = *sdlRect(&rects[i])
	}
	return &rs[0]
}

func sdlRect(rect *image.Rectangle) *C.SDL_Rect {
	if rect == nil {
		return nil
	}
	return &C.SDL_Rect{
		x: C.int(rect.Min.X),
		y: C.int(rect.Min.Y),
		w: C.int(rect.Dx()),
		h: C.int(rect.Dy()),
	}
}

// DrawPNG draws the image loaded from a PNG file to the canvas.
// The image is drawn with the upper-left corner located at x, y.
func (c Canvas) DrawPNG(path string, x, y int) {
	tex, ok := c.win.imgs[path]
	if !ok {
		img := loadPNG(path)
		tex = texture{
			tex:    texFromImage(c.win.rend, img),
			width:  img.Bounds().Dx(),
			height: img.Bounds().Dy(),
		}
		c.win.imgs[path] = tex
	}
	dst := image.Rect(x, y, x+tex.width, y+tex.height)
	if C.SDL_RenderCopy(c.win.rend, tex.tex, nil, sdlRect(&dst)) < 0 {
		panic(sdlError())
	}
}

// SetFont sets the current font face and size (in points).
func (c *Canvas) SetFont(path string, size int) {
	c.font = getFont(path, size)
}

// FillString fills a string of text in the current font and draw color.  X and y specify the
// upper-left corner of the bounding box of the text, and the width and height of the
// bounding box is returned.
func (c Canvas) FillString(s string, x, y int) (width, height int) {
	img := c.font.draw(s, c.color())
	tex := texFromImage(c.win.rend, img)
	defer C.SDL_DestroyTexture(tex)
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()
	dst := image.Rect(x, y, x+w, y+h)
	if C.SDL_RenderCopy(c.win.rend, tex, nil, sdlRect(&dst)) < 0 {
		panic(sdlError())
	}
	return w, h
}

// StringSize returns the width and height of the string in pixels when rendered in the current font.
func (c Canvas) StringSize(s string) (width, height int) {
	height, _, _ = c.font.extents()
	return c.font.width(s), height
}

func texFromImage(rend *C.SDL_Renderer, img *image.NRGBA) *C.SDL_Texture {
	b := img.Bounds()
	w, h := b.Dx(), b.Dy()
	fmt := C.SDL_PIXELFORMAT_ABGR8888
	acc := C.SDL_TEXTUREACCESS_STATIC
	tex := C.SDL_CreateTexture(rend, C.Uint32(fmt), C.int(acc), C.int(w), C.int(h))
	if tex == nil {
		panic(sdlError())
	}
	if C.SDL_UpdateTexture(tex, nil, unsafe.Pointer(&img.Pix[0]), C.int(img.Stride)) < 0 {
		panic(sdlError())
	}
	if C.SDL_SetTextureBlendMode(tex, C.SDL_BLENDMODE_BLEND) < 0 {
		panic(sdlError())
	}
	return tex
}

func loadPNG(path string) *image.NRGBA {
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
