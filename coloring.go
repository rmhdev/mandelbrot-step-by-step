package main

import (
	"errors"
	"fmt"
	"image/color"
)

type Coloring interface {
	color(v Verification) color.Color
}

func CreateColoring(name string, p Palette) (Coloring, error) {
	switch name {
	case "basic":
		return BasicColoring{p}, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined coloring '%s'", name))
}

type BasicColoring struct {
	palette Palette
}

func (c BasicColoring) color(v Verification) color.Color {
	if v.isInside {
		return c.palette.color(0)
	}
	pos := v.iterations % (c.palette.length())
	if 0 == pos {
		pos += 1
	}

	return c.palette.color(pos)
}
