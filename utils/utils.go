package utils

import (
	"encoding/json"
	"image"
	clr "image/color"
	"image/png"
	"os"

	"github.com/ath0m/DistributedRaytracer/engine"
)

// Options defines all the command line options available (all have a default value)
type Options struct {
	Width        int    // width in pixel
	Height       int    // height in pixel
	RaysPerPixel int    // number of rays per pixel
	Input        string // path to file for world definitionin in json format
	Output       string // path to file for saving
	Seed         int64  // seed for random number generator
	CPU          int    // number of CPU to use
}

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

// SaveImage saves the image (if requested) to a file in png format
func SaveImage(pixels []uint32, options Options) (error, bool) {
	if options.Output != "" {
		f, err := os.OpenFile(options.Output, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err, true
		}

		img := CreateImage(pixels, options.Width, options.Height)

		if err := png.Encode(f, img); err != nil {
			f.Close()
			return err, true
		}

		if err := f.Close(); err != nil {
			return err, true
		}

		return nil, true
	}

	return nil, false

}

func LoadWorld(file string) (*engine.World, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var world engine.World
	err = json.Unmarshal(content, &world)
	if err != nil {
		return nil, err
	}

	return &world, nil
}
