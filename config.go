package main

import (
	"errors"
	"fmt"
)

type Config struct {
	size    Size
	realMin float64
	realMax float64
	imagMin float64
	imagMax float64
}

func CreateConfig(size Size, realMin float64, realMax float64, imagMin float64) Config {
	imagMax := imagMin + (realMax-realMin)*float64(size.rawHeight())/float64(size.rawWidth())

	return Config{size, realMin, realMax, imagMin, imagMax}
}

func (c Config) toReal(x int) (float64, error) {
	if (x < 0) || x >= c.size.rawWidth() {
		return 0, errors.New(fmt.Sprintf("X is out of bounds! Got %d, expected [0..%d]", x, c.size.rawWidth()-1))
	}
	size := ((c.realMax - c.realMin) / float64(c.size.rawWidth()-1))

	return c.realMin + float64(x)*size, nil
}

func (c Config) toImag(y int) (float64, error) {
	if (y < 0) || (y >= c.size.rawHeight()) {
		return 0, errors.New(fmt.Sprintf("Y is out of bounds! Got: %d, expected: [0..%d]", y, c.size.rawHeight()-1))
	}
	size := ((c.imagMax - c.imagMin) / float64(c.size.rawHeight()-1))

	return c.imagMax - float64(y)*size, nil
}

func (c Config) representation(verifier Verifier) Representation {
	representation := CreateRepresentation(c.size)
	realC, imagC := 0.0, 0.0
	for y := 0; y < c.size.rawHeight(); y++ {
		imagC, _ = c.toImag(y)
		for x := 0; x < c.size.rawWidth(); x++ {
			realC, _ = c.toReal(x)
			representation.set(x, y, verifier.verify(realC, imagC))
		}
	}
	return representation
}
