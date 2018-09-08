package main

import "fmt"

func main() {
	config := Config{101, 41, -2.0, 0.5, -1.0, 1.0}
	verifier := Verifier{50}
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
