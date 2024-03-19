package engine

import (
	"image"
	clr "image/color"
)

func CreateImage(pixels []uint32, width int, height int) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	k := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := pixels[k]
			img.Set(x, y, clr.NRGBA{
				R: uint8(p >> 16 & 0xFF),
				G: uint8(p >> 8 & 0xFF),
				B: uint8(p & 0xFF),
				A: 255,
			})
			k++
		}
	}

	return img
}
