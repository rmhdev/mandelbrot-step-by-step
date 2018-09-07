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

func (c Config) toReal(x int) (float64, error) {
	if (x < 0) || x >= c.width {
		return 0, errors.New(fmt.Sprintf("X is out of bounds! Got %d, expected [0..%d]", x, c.width-1))
	}
	size := ((c.realMax - c.realMin) / float64(c.width-1))

	return c.realMin + float64(x)*size, nil
}
