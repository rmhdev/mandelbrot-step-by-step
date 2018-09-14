package main

import "fmt"

type Exporter struct {
	representation Representation
}

func (e Exporter) export() string {
	result := ""
	for y := 0; y < e.representation.height(); y++ {
		line := ""
		for x := 0; x < e.representation.width(); x++ {
			if e.representation.isInside(x, y) {
				line += "*"
			} else {
				line += "·"
			}
		}
		result += fmt.Sprintln(line)
	}

	return result
}
