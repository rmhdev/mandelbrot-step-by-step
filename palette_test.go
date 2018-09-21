package main

import (
	"image/color"
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
