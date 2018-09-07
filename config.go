package main

type Config struct {
	width   int
	height  int
	realMin float64
	realMax float64
	imagMin float64
	imagMax float64
}

func (c Config) toReal(x int) float64 {
	return c.realMin
}
