package main

import (
	"image/color"
	"testing"
)

func TestBasicColoring(t *testing.T) {
	palette := BobRossPalette{}
	tests := []struct {
		isInside   bool
		iterations int
		expected   color.Color
	}{
		{true, 1, palette.color(0)},
		{false, 1, palette.color(1)},
		{false, palette.length() - 1, palette.color(palette.length() - 1)},
		{false, palette.length(), palette.color(1)},
	}
	coloring := Coloring{palette}
	for _, test := range tests {
		v := Verification{test.isInside, test.iterations}
		result := coloring.color(v)
		if test.expected != result {
			t.Errorf("Incorrect basic color(inInside: %v, iter: %d). Got: %v, expected: %v", test.isInside, test.iterations, result, test.expected)
		}
	}
}
