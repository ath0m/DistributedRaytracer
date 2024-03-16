package engine

import (
	"math"
)

// Color defines the basic Red/Green/Blue as raw float64 values
type Color struct {
	R, G, B float64
}

var (
	White = Color{1.0, 1.0, 1.0}
	Black = Color{}
)

// Scale scales the Color by the value (return a new Color)
func (c Color) Scale(t float64) Color {
	return Color{R: c.R * t, G: c.G * t, B: c.B * t}
}

// Mult Multiplies 2 colors together (component by component multiplication)
func (c Color) Mult(c2 Color) Color {
	return Color{R: c.R * c2.R, G: c.G * c2.G, B: c.B * c2.B}
}

// Add adds the 2 colors (return a new color)
func (c Color) Add(c2 Color) Color {
	return Color{R: c.R + c2.R, G: c.G + c2.G, B: c.B + c2.B}
}

// PixelValue converts a raw Color into a pixel value (0-255) packed into a uint32
func (c Color) PixelValue() uint32 {
	r := uint32(math.Min(255.0, c.R*255.99))
	g := uint32(math.Min(255.0, c.G*255.99))
	b := uint32(math.Min(255.0, c.B*255.99))

	return ((r & 0xFF) << 16) | ((g & 0xFF) << 8) | (b & 0xFF)
}
