package main

import "image/color"

type Coloring struct {
	palette Palette
}

func (c Coloring) color(v Verification) color.Color {
	if v.isInside {
		return c.palette.color(0)
	}
	pos := v.iterations % (c.palette.length())
	if 0 == pos {
		pos += 1
	}

	return c.palette.color(pos)
}
