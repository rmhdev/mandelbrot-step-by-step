package main

import (
	"image/color"
	"reflect"
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
	coloring := BasicColoring{palette}
	for _, test := range tests {
		v := Verification{test.isInside, test.iterations, 0.0, 0.0}
		result := coloring.color(v)
		if test.expected != result {
			t.Errorf("Incorrect basic color(inInside: %v, iter: %d). Got: %v, expected: %v", test.isInside, test.iterations, result, test.expected)
		}
	}
}

func TestCreateCorrectColoring(t *testing.T) {
	palette, _ := CreatePalette("bob_ross")
	tests := []struct {
		name     string
		expected Coloring
	}{
		{"basic", BasicColoring{palette}},
	}

	for _, test := range tests {
		coloring, err := CreateColoring(test.name, palette)
		if err != nil {
			t.Errorf("Incorrect '%s' coloring creation, got error", test.name)
		}
		if reflect.TypeOf(coloring) != reflect.TypeOf(test.expected) {
			t.Errorf("Incorrect '%s' coloring type, got: '%v', expected: '%v'", test.name, reflect.TypeOf(coloring), reflect.TypeOf(test.expected))
		}
	}
}

func TestCreateIncorrectColoring(t *testing.T) {
	palette, _ := CreatePalette("bob_ross")
	coloring, err := CreateColoring("lorem", palette)
	if err == nil {
		t.Errorf("Incorrect '%s' coloring creation, expected error", "lorem")
	}
	if coloring != nil {
		t.Errorf("Incorrect '%s' coloring type, got: '%v', expected: nil", "lorem", reflect.TypeOf(coloring))
	}
}
