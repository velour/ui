package sdl2

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
	"unsafe"
)

type Renderer C.SDL_Renderer

func (rend *Renderer) sdl() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(rend)
}

type RendererFlags C.Uint32

const (
	// Software says the renderer is a software fallback.
	Software RendererFlags = C.SDL_RENDERER_SOFTWARE

	// Accelerated says the renderer uses hardware acceleration.
	Accelerated RendererFlags = C.SDL_RENDERER_ACCELERATED

	// PresentVsync says present is synchronized with the refresh rate.
	PresentVsync RendererFlags = C.SDL_RENDERER_PRESENTVSYNC

	// TargetTexture says the renderer supports rendering to texture.
	TargetTexture RendererFlags = C.SDL_RENDERER_TARGETTEXTURE
)

// CreateRenderer creates a 2D rendering context for a window.
func CreateRenderer(win *Window, index int, flags RendererFlags) (*Renderer, error) {
	rend := C.SDL_CreateRenderer(win.sdl(), C.int(index), C.Uint32(flags))
	if rend == nil {
		return nil, sdlError()
	}
	return (*Renderer)(rend), nil
}

// Destroy destroys the rendering context for a window and free associated textures.
func (rend *Renderer) Destroy() {
	C.SDL_DestroyRenderer(rend.sdl())
}

// Present updates the screen with rendering performed.
func (rend *Renderer) Present() {
	C.SDL_RenderPresent(rend.sdl())
}

// SetDrawColor sets the color used for drawing operations (Rect, Line and Clear).
func (rend *Renderer) SetDrawColor(col color.Color) error {
	r, g, b, a := col.RGBA()
	f := 255.0 / 0xFFFF
	r8 := C.Uint8(float64(r) * f)
	g8 := C.Uint8(float64(g) * f)
	b8 := C.Uint8(float64(b) * f)
	a8 := C.Uint8(float64(a) * f)
	if C.SDL_SetRenderDrawColor(rend.sdl(), r8, g8, b8, a8) < 0 {
		return sdlError()
	}
	return nil
}

type BlendMode C.SDL_BlendMode

const (
	// None is no blending.
	None BlendMode = C.SDL_BLENDMODE_NONE

	// Blend computes dst = (src * A) + (dst * (1-A))
	Blend BlendMode = C.SDL_BLENDMODE_BLEND

	// Add computes dst = (src * A) + dst.
	Add BlendMode = C.SDL_BLENDMODE_ADD

	// Mod computes dst = src * dst
	Mod BlendMode = C.SDL_BLENDMODE_MOD
)

// SetDrawBlendMode sets the blend mode used for drawing operations (Fill and Line).
func (rend *Renderer) SetDrawBlendMode(mode BlendMode) error {
	if C.SDL_SetRenderDrawBlendMode(rend.sdl(), C.SDL_BlendMode(mode)) < 0 {
		return sdlError()
	}
	return nil
}

// DrawPoint draws a point on the current rendering target.
func (rend *Renderer) DrawPoint(x, y int) error {
	if C.SDL_RenderDrawPoint(rend.sdl(), C.int(x), C.int(y)) < 0 {
		return sdlError()
	}
	return nil
}

// DrawPoints draws multiple points on the current rendering target.
func (rend *Renderer) DrawPoints(points []image.Point) error {
	pts := make([]C.SDL_Point, len(points))
	for i, p := range points {
		pts[i].x = C.int(p.X)
		pts[i].y = C.int(p.Y)
	}
	if C.SDL_RenderDrawPoints(rend.sdl(), &pts[0], C.int(len(points))) < 0 {
		return sdlError()
	}
	return nil
}

// DrawLine draws a line on the current rendering target.
func (rend *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	if C.SDL_RenderDrawLine(rend.sdl(), C.int(x1), C.int(y1), C.int(x2), C.int(y2)) < 0 {
		return sdlError()
	}
	return nil
}

// DrawLines draws a series of connected lines on the current rendering target.
func (rend *Renderer) DrawLines(points []image.Point) error {
	pts := make([]C.SDL_Point, len(points))
	for i, p := range points {
		pts[i].x = C.int(p.X)
		pts[i].y = C.int(p.Y)
	}
	if C.SDL_RenderDrawLines(rend.sdl(), &pts[0], C.int(len(points))) < 0 {
		return sdlError()
	}
	return nil
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

// FillRect fills a rectangle on the current rendering target with the drawing color.
func (rend *Renderer) FillRect(rect *image.Rectangle) error {
	if C.SDL_RenderFillRect(rend.sdl(), sdlRect(rect)) < 0 {
		return sdlError()
	}
	return nil
}

// FillRects fills some number of rectangles on the current rendering target with the drawing color.
func (rend *Renderer) FillRects(rects []image.Rectangle) error {
	rs := make([]C.SDL_Rect, len(rects))
	for i := range rects {
		rs[i] = *sdlRect(&rects[i])
	}
	if C.SDL_RenderFillRects(rend.sdl(), &rs[0], C.int(len(rects))) < 0 {
		return sdlError()
	}
	return nil
}

// DratRect draws a rectangle on the current rendering target.
func (rend *Renderer) DrawRect(rect *image.Rectangle) error {
	if C.SDL_RenderDrawRect(rend.sdl(), sdlRect(rect)) < 0 {
		return sdlError()
	}
	return nil
}

// DrawRects draws some number of rectangles on the current rendering target.
func (rend *Renderer) DrawRects(rects []image.Rectangle) error {
	rs := make([]C.SDL_Rect, len(rects))
	for i := range rects {
		rs[i] = *sdlRect(&rects[i])
	}
	if C.SDL_RenderDrawRects(rend.sdl(), &rs[0], C.int(len(rects))) < 0 {
		return sdlError()
	}
	return nil
}

// Copy copies a portion of the texture to the current rendering target.
func (rend *Renderer) Copy(tex *Texture, src, dst *image.Rectangle) error {
	if C.SDL_RenderCopy(rend.sdl(), tex.sdl(), sdlRect(src), sdlRect(dst)) < 0 {
		return sdlError()
	}
	return nil
}

type Texture C.SDL_Texture

func (tex *Texture) sdl() *C.SDL_Texture {
	return (*C.SDL_Texture)(tex)
}

type PixelFormat C.Uint32

const (
	RGBA8888 PixelFormat = C.SDL_PIXELFORMAT_RGBA8888
	ABGR8888 PixelFormat = C.SDL_PIXELFORMAT_ABGR8888
)

type TextureAccess C.SDL_TextureAccess

const (
	// Static textures change rarely and are not lockable.
	Static TextureAccess = C.SDL_TEXTUREACCESS_STATIC

	// Streaming textures change frequently and are lockable.
	Streaming TextureAccess = C.SDL_TEXTUREACCESS_STREAMING
)

// CreateTexture creates a texture for a rendering context.
func CreateTexture(rend *Renderer, fmt PixelFormat, acc TextureAccess, w, h int) (*Texture, error) {
	tex := C.SDL_CreateTexture(rend.sdl(),
		C.Uint32(fmt),
		C.int(acc),
		C.int(w),
		C.int(h))
	if tex == nil {
		return nil, sdlError()
	}
	return (*Texture)(tex), nil
}

// CreateTextureFromImage creates a texture from an existing image.
func CreateTextureFromImage(rend *Renderer, img *image.NRGBA) (*Texture, error) {
	b := img.Bounds()
	tex, err := CreateTexture(rend, RGBA8888, Static, b.Dx(), b.Dy())
	if err != nil {
		return nil, err
	}
	if err := tex.Update(nil, img.Pix, img.Stride); err != nil {
		tex.Destroy()
		return nil, err
	}
	return tex, nil
}

// Destroy destroys the specified texture.
func (tex *Texture) Destroy() {
	C.SDL_DestroyTexture(tex.sdl())
}

// Update updates the given texture rectangle with new pixel data.
func (tex *Texture) Update(rect *image.Rectangle, data []byte, pitch int) error {
	if C.SDL_UpdateTexture(tex.sdl(), sdlRect(rect), unsafe.Pointer(&data[0]), C.int(pitch)) < 0 {
		return sdlError()
	}
	return nil
}

func (tex *Texture) SetBlendMode(mode BlendMode) error {
	if C.SDL_SetTextureBlendMode(tex.sdl(), C.SDL_BlendMode(mode)) < 0 {
		return sdlError()
	}
	return nil
}
