// © 2013 the Ui Authors under the MIT license. See AUTHORS for the list of authors.

package ui

/*
#include "ui.h"
*/
import "C"

import (
	"image"
	"image/color"
	"io/ioutil"
	"os"

	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
)

const (
	// PxInch is the number of pixels in an inch.
	// 72 seems to make font sizes consistent with SDL_ttf…
	pxInch = 72.0

	// PtInch is the number of Postscript points in an inch.
	ptInch = 72.0
)

var fonts = make(map[string]font)

type font struct {
	size int
	path string
	*truetype.Font
	*freetype.Context
}

func getFont(path string, sizePts int) font {
	sizePx := int(float64(sizePts)/ptInch*pxInch + 0.5)

	if f, ok := fonts[path]; ok {
		f.size = sizePx
		return f
	}

	in, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	fdata, err := ioutil.ReadAll(in)
	if err != nil {
		panic(err)
	}

	f := font{
		path:    path,
		Context: freetype.NewContext(),
	}
	if f.Font, err = truetype.Parse(fdata); err != nil {
		panic(err)
	}
	f.SetFont(f.Font)
	f.SetDPI(pxInch)
	fonts[path] = f

	f.size = sizePx
	return f
}

func (f *font) extents() (height, ascent, descent int) {
	scale := f.scale()
	bounds := f.Bounds(f.FUnitsPerEm())
	ascent = int(float64(bounds.YMax)*scale + 0.5)
	descent = int(float64(bounds.YMin)*scale - 0.5)
	return ascent - descent, ascent, descent
}

func (f *font) width(s string) int {
	em := f.FUnitsPerEm()
	var width int32
	prev, hasPrev := truetype.Index(0), false
	for _, r := range s {
		index := f.Index(r)
		if hasPrev {
			width += f.Kerning(em, prev, index)
		}
		width += f.HMetric(em, index).AdvanceWidth
		prev, hasPrev = index, true
	}
	return int(float64(width)*f.scale() + 0.5)
}

func (f *font) scale() float64 {
	em := f.FUnitsPerEm()
	return (float64(f.size) / ptInch * pxInch) / float64(em)
}

func (f *font) draw(s string, col color.Color) *image.NRGBA {
	width := f.width(s)
	height, _, descent := f.extents()
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	f.SetFontSize(float64(f.size))
	f.SetSrc(image.NewUniform(col))
	f.SetClip(img.Bounds())
	f.SetDst(img)
	if _, err := f.DrawString(s, freetype.Pt(0, height+descent)); err != nil {
		panic(err)
	}
	return img
}
