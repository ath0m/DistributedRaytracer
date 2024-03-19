package color

import (
	"math"
	"testing"
)

func equalColor(c1, c2 Color) bool {
	const epsilon = 1e-8
	return math.Abs(c1.R-c2.R) < epsilon && math.Abs(c1.G-c2.G) < epsilon && math.Abs(c1.B-c2.B) < epsilon
}

func TestColorScale(t *testing.T) {
	cases := []struct {
		color    Color
		scale    float64
		expected Color
	}{
		{Color{0.0, 0.0, 0.0}, 2.0, Color{0.0, 0.0, 0.0}},
		{Color{1.0, 1.0, 1.0}, 0.5, Color{0.5, 0.5, 0.5}},
		{Color{0.5, 0.5, 0.5}, 1.5, Color{0.75, 0.75, 0.75}},
		{Color{0.25, 0.75, 0.5}, 0.1, Color{0.025, 0.075, 0.05}},
	}
	for _, tc := range cases {
		result := tc.color.Scale(tc.scale)
		if !equalColor(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestColorMult(t *testing.T) {
	cases := []struct {
		color1   Color
		color2   Color
		expected Color
	}{
		{Color{0.0, 0.0, 0.0}, Color{0.0, 0.0, 0.0}, Color{0.0, 0.0, 0.0}},
		{Color{1.0, 1.0, 1.0}, Color{1.0, 1.0, 1.0}, Color{1.0, 1.0, 1.0}},
		{Color{0.5, 0.5, 0.5}, Color{0.5, 0.5, 0.5}, Color{0.25, 0.25, 0.25}},
		{Color{0.25, 0.75, 0.5}, Color{0.5, 0.25, 0.75}, Color{0.125, 0.1875, 0.375}},
	}

	for _, tc := range cases {
		result := tc.color1.Mult(tc.color2)
		if !equalColor(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestColorAdd(t *testing.T) {
	cases := []struct {
		color1   Color
		color2   Color
		expected Color
	}{
		{Color{0.0, 0.0, 0.0}, Color{0.0, 0.0, 0.0}, Color{0.0, 0.0, 0.0}},
		{Color{1.0, 1.0, 1.0}, Color{1.0, 1.0, 1.0}, Color{2.0, 2.0, 2.0}},
		{Color{0.5, 0.5, 0.5}, Color{0.5, 0.5, 0.5}, Color{1.0, 1.0, 1.0}},
		{Color{0.25, 0.75, 0.5}, Color{0.25, 0.75, 0.5}, Color{0.5, 1.5, 1.0}},
	}

	for _, tc := range cases {
		result := tc.color1.Add(tc.color2)
		if !equalColor(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}

func TestColorPixelValue(t *testing.T) {
	cases := []struct {
		color    Color
		expected uint32
	}{
		{Color{0.0, 0.0, 0.0}, 0x000000},
		{Color{1.0, 1.0, 1.0}, 0xFFFFFF},
		{Color{0.5, 0.5, 0.5}, 0x7F7F7F},
		{Color{0.25, 0.75, 0.5}, 0x3FBF7F},
	}

	for _, tc := range cases {
		result := tc.color.PixelValue()
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
