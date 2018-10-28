package main

import (
	"errors"
	"fmt"
)

const (
	MaxWidth  int = 5000
	MaxHeight int = 5000
	MaxFactor int = 16
)

type Size struct {
	width  int
	height int
	factor int
}

func CreateSize(width int, height int, factor int) (Size, error) {
	if 0 >= width || MaxWidth < width {
		return Size{1, 1, 1}, errors.New(fmt.Sprintf("Width is out of bounds! Got %d, expected [1..%d]", width, MaxWidth))
	}
	if 0 >= height || MaxHeight < height {
		return Size{1, 1, 1}, errors.New(fmt.Sprintf("Height is out of bounds! Got %d, expected [1..%d]", height, MaxHeight))
	}
	if 0 >= factor || MaxFactor < factor {
		return Size{1, 1, 1}, errors.New(fmt.Sprintf("Detail is out of bounds! Got %d, expected [1..%d]", factor, MaxFactor))
	}
	return Size{width, height, factor}, nil
}

func (s Size) rawWidth() int {
	return s.width * s.factor
}

func (s Size) rawHeight() int {
	return s.height * s.factor
}
