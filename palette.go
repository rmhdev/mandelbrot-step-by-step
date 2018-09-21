package main

import "image/color"

type BlackWhitePalette struct {
}

func (p BlackWhitePalette) color(v Verification) color.Color {
	if v.isInside {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{255, 255, 255, 255}
}
