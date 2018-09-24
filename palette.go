package main

import (
	"errors"
	"fmt"
	"image/color"
)

type Palette interface {
	color(v Verification) color.Color
}

func CreatePalette(name string) (Palette, error) {
	switch name {
	case "bw":
		return BlackWhitePalette{}, nil
	case "bob_ross":
		return BobRossPalette{}, nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined palette '%s'", name))
}

type BlackWhitePalette struct {
}

func (p BlackWhitePalette) color(v Verification) color.Color {
	if v.isInside {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{255, 255, 255, 255}
}

var BobRoss = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // Midnight black
	color.RGBA{0x02, 0x1e, 0x44, 0xff}, // Prussian blue
	color.RGBA{0x0a, 0x34, 0x10, 0xff}, // Sap green
	color.RGBA{0x0c, 0x00, 0x40, 0xff}, // Phthalo blue
	color.RGBA{0x10, 0x2e, 0x3c, 0xff}, // Phthalo green
	color.RGBA{0x22, 0x1b, 0x15, 0xff}, // Van Dyke brown
	color.RGBA{0x4e, 0x15, 0x00, 0xff}, // Alizarin crimson
	color.RGBA{0x5f, 0x2e, 0x1f, 0xff}, // Dark Sienna
	color.RGBA{0xc7, 0x9b, 0x00, 0xff}, // Yellow ochre
	color.RGBA{0xdb, 0x00, 0x00, 0xff}, // Bright red
	color.RGBA{0xff, 0x3c, 0x00, 0xff}, // Cadmium yellow
	color.RGBA{0xff, 0xb8, 0x00, 0xff}, // Indian yellow
	color.RGBA{0xff, 0xff, 0xff, 0xff}, // Titanium white
}

type BobRossPalette struct {
}

func (p BobRossPalette) color(v Verification) color.Color {
	if v.isInside {
		return BobRoss[0]
	}
	pos := (v.iterations) % (len(BobRoss) - 1)

	return BobRoss[pos+1]
}
