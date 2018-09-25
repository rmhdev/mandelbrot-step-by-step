package main

import (
	"image/color"
	"reflect"
	"testing"
)

func TestBlackWhitePalette(t *testing.T) {
	palette := BlackWhitePalette{}
	expectedLength := 2
	if expectedLength != palette.length() {
		t.Errorf("Incorrect length in BlackWhite palette. Got: '%d', expected: '%d'", palette.length(), expectedLength)
	}

	tests := []struct {
		position int
		expected color.Color
	}{
		{0, color.RGBA{0, 0, 0, 255}},
		{1, color.RGBA{255, 255, 255, 255}},
		{2, color.RGBA{0, 0, 0, 255}},
		{-1, color.RGBA{255, 255, 255, 255}},
	}
	for _, test := range tests {
		result := palette.color(test.position)
		if test.expected != result {
			t.Errorf("Incorrect color in BW palette on pos(%d). Got: '%v', expected: '%v'", test.position, result, test.expected)
		}
	}
}

func TestBobRossPalette(t *testing.T) {
	palette := BobRossPalette{}
	expectedLength := 13
	if expectedLength != palette.length() {
		t.Errorf("Incorrect length in BobRoss palette. Got: '%d', expected: '%d'", palette.length(), expectedLength)
	}

	tests := []struct {
		position    int
		expected    color.Color
		description string
	}{
		{0, color.RGBA{0, 0, 0, 255}, "First color"},
		{12, color.RGBA{255, 255, 255, 255}, "Last color"},
		{13, color.RGBA{0, 0, 0, 255}, "Position is greater than number of colors"},
		{-1, color.RGBA{2, 30, 68, 255}, "Use absolute values"},
	}
	for _, test := range tests {
		result := palette.color(test.position)
		if test.expected != result {
			t.Errorf("Incorrect color in BobRoss palette (position: %d). Got: '%v', expected: '%v' (%s)", test.position, result, test.expected, test.description)
		}
	}
}

func TestPlan9Palette(t *testing.T) {
	palette := Plan9Palette{}
	expectedLength := 256
	if expectedLength != palette.length() {
		t.Errorf("Incorrect length in BlackWhite palette. Got: '%d', expected: '%d'", palette.length(), expectedLength)
	}

	tests := []struct {
		position    int
		expected    color.Color
		description string
	}{
		{0, color.RGBA{0, 0, 0, 255}, "First color"},
		{255, color.RGBA{255, 255, 255, 255}, "Last color"},
		{256, color.RGBA{0, 0, 0, 255}, "Position is greater than number of colors"},
		{-1, color.RGBA{0, 0, 68, 255}, "use absolute values"},
	}
	for _, test := range tests {
		result := palette.color(test.position)
		if test.expected != result {
			t.Errorf("Incorrect color in Plan9 palette (position: %d). Got: '%v', expected: '%v' (%s)", test.position, result, test.expected, test.description)
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
