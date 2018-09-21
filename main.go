package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	width := flag.Int("width", 804, "width")
	height := flag.Int("height", 603, "height")
	realMin := flag.Float64("realMin", -2.5, "Minimum value for real part")
	realMax := flag.Float64("realMax", 1.0, "Maximum value for real part")
	imagMin := flag.Float64("imagMin", -1.3125, "Minimum value for imaginary part")
	iterations := flag.Int("iterations", 50, "maximum number of iterations per pixel")
	exporterName := flag.String("exporter", "image", "name of the exporter")
	folder := flag.String("folder", "mandelbrot", "folder for exporting images")
	filename := flag.String("filename", "", "name of the image")

	flag.Parse() // Don't forget this!

	config := CreateConfig(*width, *height, *realMin, *realMax, *imagMin)
	representation := config.representation(Verifier{*iterations})
	palette := BlackWhitePalette{}
	exporter, exporterErr := CreateExporter(*exporterName, representation, *folder, *filename, palette)
	if exporterErr != nil {
		fmt.Print(exporterErr)
		os.Exit(1)
	}
	result, err := exporter.export()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Print(result)
}
