package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	width := flag.Int("width", 800, "width")
	height := flag.Int("height", 601, "height")
	realMin := flag.Float64("realMin", -2.0, "Minimum value for real part")
	realMax := flag.Float64("realMax", 0.5, "Maximum value for real part")
	imagMin := flag.Float64("imagMin", -1.0, "Minimum value for imaginary part")
	imagMax := flag.Float64("imagMax", 1.0, "Maximum value for imaginary part")
	iterations := flag.Int("iterations", 50, "maximum number of iterations per pixel")
	exporterName := flag.String("exporter", "image", "name of the exporter")
	folder := flag.String("folder", "mandelbrot", "folder for exporting images")
	filename := flag.String("filename", "", "name of the image")

	flag.Parse() // Don't forget this!

	config := Config{*width, *height, *realMin, *realMax, *imagMin, *imagMax}
	representation := config.representation(Verifier{*iterations})
	exporter, exporterErr := CreateExporter(*exporterName, representation, *folder, *filename)
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
