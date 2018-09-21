package main

import (
	"errors"
	"fmt"
)

type Config struct {
	width   int
	height  int
	realMin float64
	realMax float64
	imagMin float64
	imagMax float64
}

func CreateConfig(width int, height int, realMin float64, realMax float64, imagMin float64) Config {
	imagMax := imagMin + (realMax-realMin)*float64(height)/float64(width)

	return Config{width, height, realMin, realMax, imagMin, imagMax}
}

func (c Config) toReal(x int) (float64, error) {
	if (x < 0) || x >= c.width {
		return 0, errors.New(fmt.Sprintf("X is out of bounds! Got %d, expected [0..%d]", x, c.width-1))
	}
	size := ((c.realMax - c.realMin) / float64(c.width-1))

	return c.realMin + float64(x)*size, nil
}

func (c Config) toImag(y int) (float64, error) {
	if (y < 0) || (y >= c.height) {
		return 0, errors.New(fmt.Sprintf("Y is out of bounds! Got: %d, expected: [0..%d]", y, c.height-1))
	}
	size := ((c.imagMax - c.imagMin) / float64(c.height-1))

	return c.imagMax - float64(y)*size, nil
}

func (c Config) representation(verifier Verifier) Representation {
	representation := CreateRepresentation(c.width, c.height)
	realC, imagC := 0.0, 0.0
	for y := 0; y < c.height; y++ {
		imagC, _ = c.toImag(y)
		for x := 0; x < c.width; x++ {
			realC, _ = c.toReal(x)
			representation.set(x, y, verifier.verify(realC, imagC))
		}
	}
	return representation
}
