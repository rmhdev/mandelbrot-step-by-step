package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestBlackWhitePalette(t *testing.T) {
	palette := BlackWhitePalette{}

	black := color.RGBA{0, 0, 0, 255}
	isInsideColor := palette.color(Verification{true, 1})
	if black != isInsideColor {
		t.Errorf("Incorrect color in BW palette. Got: '%v', expected: '%v'", isInsideColor, black)
	}

	white := color.RGBA{255, 255, 255, 255}
	isNotInsideColor := palette.color(Verification{false, 1})
	if white != isNotInsideColor {
		t.Errorf("Incorrect color in BW palette. Got: '%v', expected: '%v'", isNotInsideColor, white)
	}
}

func TestBobRossPalette(t *testing.T) {
	palette := BobRossPalette{}
	tests := []struct {
		isInside      bool
		iterations    int
		expectedColor color.Color
		description   string
	}{
		{true, 1, color.RGBA{0, 0, 0, 255}, "Point is inside Mandelbrot"},
		{false, 11, color.RGBA{255, 255, 255, 255}, "Max iteration, and still not inside"},
		{false, 12, color.RGBA{2, 30, 68, 255}, "Iteration greater than number of colors"},
	}
	for _, test := range tests {
		resultColor := palette.color(Verification{test.isInside, test.iterations})
		if test.expectedColor != resultColor {
			t.Errorf("Incorrect color in BobRoss palette (inside: %v, iter: %d). Got: '%v', expected: '%v' (%s)", test.isInside, test.iterations, resultColor, test.expectedColor, test.description)
		}
	}
}

func TestPlan9Palette(t *testing.T) {
	palette := Plan9Palette{}
	tests := []struct {
		isInside      bool
		iterations    int
		expectedColor color.Color
		description   string
	}{
		{true, 1, color.RGBA{0, 0, 0, 255}, "Point is inside Mandelbrot"},
		{false, 254, color.RGBA{255, 255, 255, 255}, "Max iteration, and still not inside"},
		{false, 255, color.RGBA{0, 0, 68, 255}, "Iteration greater than number of colors"},
	}
	for _, test := range tests {
		resultColor := palette.color(Verification{test.isInside, test.iterations})
		if test.expectedColor != resultColor {
			t.Errorf("Incorrect color in Plan9 palette (inside: %v, iter: %d). Got: '%v', expected: '%v' (%s)", test.isInside, test.iterations, resultColor, test.expectedColor, test.description)
		}
	}
}

func TestCreateCorrectPalette(t *testing.T) {
	tests := []struct {
		name     string
		expected Palette
	}{
		{"bw", BlackWhitePalette{}},
		{"bob_ross", BobRossPalette{}},
		{"plan9", Plan9Palette{}},
	}

	for _, test := range tests {
		palette, err := CreatePalette(test.name)
		if err != nil {
			t.Errorf("Incorrect '%s' palette creation, got error", test.name)
		}
		if reflect.TypeOf(palette) != reflect.TypeOf(test.expected) {
			t.Errorf("Incorrect '%s' palette type, got: '%v', expected: '%v'", test.name, reflect.TypeOf(palette), reflect.TypeOf(test.expected))
		}
	}
}

func TestCreateIncorrectPalette(t *testing.T) {
	palette, err := CreatePalette("lorem")
	if err == nil {
		t.Errorf("Incorrect '%s' palette creation, expected error", "lorem")
	}
	if palette != nil {
		t.Errorf("Incorrect '%s' palette type, got: '%v', expected: nil", "lorem", reflect.TypeOf(palette))
	}
}
