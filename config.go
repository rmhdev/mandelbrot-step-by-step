package main

import (
	"errors"
	"fmt"
	"math"
)

type Config struct {
	size       Size
	iterations int
	realMin    float64
	realMax    float64
	imagMin    float64
	imagMax    float64
}

func CreateConfig(size Size, iterations int, realMin float64, realMax float64, imagMin float64) Config {
	imagMax := imagMin + (realMax-realMin)*float64(size.rawHeight())/float64(size.rawWidth())

	return Config{size, iterations, realMin, realMax, imagMin, imagMax}
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

func (c Config) representation(p Progress, filename string) Representation {
	verifier := c.verifier()
	representation := CreateRepresentation(c)
	realC, imagC := 0.0, 0.0
	progressText := filename + ". build "
	total := c.size.rawHeight()
	previous := 0
	actual := 0
	for y := 0; y < c.size.rawHeight(); y++ {
		imagC, _ = c.toImag(y)
		for x := 0; x < c.size.rawWidth(); x++ {
			realC, _ = c.toReal(x)
			representation.set(x, y, verifier.verify(realC, imagC))
		}
		actual = int(p.maxBars * (y + 1) / total)
		if actual > previous {
			p.bar(progressText, p.maxBars, actual)
			previous = actual
		}
	}
	p.writeln("ok")

	return representation
}

func (c Config) verifier() Verifier {
	return Verifier{c.iterations}
}

func (c Config) center() (float64, float64) {
	realRadius, imagRadius := c.radius()

	return c.realMin + realRadius, c.imagMin + imagRadius
}

func (c Config) radius() (float64, float64) {
	realRadius := c.realMax - (c.realMax - math.Abs(c.realMax-c.realMin)/2)
	imagRadius := c.imagMax - (c.imagMax - math.Abs(c.imagMax-c.imagMin)/2)

	return realRadius, imagRadius
}
