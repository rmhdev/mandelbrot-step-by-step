package main

import (
	"flag"
	"fmt"
)

func main() {
	width := flag.Int("width", 101, "width")
	height := flag.Int("height", 41, "height")
	realMin := flag.Float64("realMin", -2.0, "Minimum value for real part")
	realMax := flag.Float64("realMax", 0.5, "Maximum value for real part")
	imagMin := flag.Float64("imagMin", -1.0, "Minimum value for imaginary part")
	imagMax := flag.Float64("imagMax", 1.0, "Maximum value for imaginary part")
	iterations := flag.Int("iterations", 50, "maximum number of iterations per pixel")

	flag.Parse() // Don't forget this!

	config := Config{*width, *height, *realMin, *realMax, *imagMin, *imagMax}
	verifier := Verifier{*iterations}

	realC, imagC := 0.0, 0.0
	for y := 0; y < config.height; y++ {
		imagC, _ = config.toImag(y)
		for x := 0; x < config.width; x++ {
			realC, _ = config.toReal(x)
			if verifier.isInside(realC, imagC) {
				fmt.Print("*")
			} else {
				fmt.Print("·")
			}
		}
		fmt.Println("")
	}
}
