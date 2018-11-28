package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	width := flag.Int("width", 804, "width")
	height := flag.Int("height", 603, "height")
	factor := flag.Int("aa", 1, "antialiasing factor")
	realMin := flag.Float64("realMin", -2.5, "minimum value for real part")
	realMax := flag.Float64("realMax", 1.0, "maximum value for real part")
	imagMin := flag.Float64("imagMin", -1.3125, "minimum value for imaginary part")
	iterations := flag.Int("iterations", 50, "maximum number of iterations per pixel")

	zoomRealCenter := flag.Float64("zoomRealCenter", -1.6, "zoom center for real part")
	zoomImagCenter := flag.Float64("zoomImagCenter", 0.0, "zoom center for imag part")
	zoomRatio := flag.Float64("zoomRatio", 0.1, "zoom ratio")
	zoomSteps := flag.Int("zoomSteps", 1, "zoom steps")

	exporterName := flag.String("exporter", "image", "name of the exporter")
	folder := flag.String("folder", "mandelbrot", "folder for exporting images")
	filename := flag.String("filename", "", "name of the image")
	paletteName := flag.String("palette", "bw", "name of the color palette")
	coloringName := flag.String("coloring", "basic", "name of the coloring method")

	flag.Parse() // Don't forget this!

	size, sizeErr := CreateSize(*width, *height, *factor)
	if sizeErr != nil {
		fmt.Print(sizeErr)
		os.Exit(1)
	}
	palette, paletteErr := CreatePalette(*paletteName)
	if paletteErr != nil {
		fmt.Print(paletteErr)
		os.Exit(1)
	}
	coloring, coloringErr := CreateColoring(*coloringName, palette)
	if coloringErr != nil {
		fmt.Print(coloringErr)
		os.Exit(1)
	}

	zoom, zoomErr := CreateZoom(*zoomSteps, *zoomRatio, *zoomRealCenter, *zoomImagCenter)
	if zoomErr != nil {
		fmt.Print(zoomErr)
		os.Exit(1)
	}

	fileNameValue := *filename
	config := CreateConfig(size, *iterations, *realMin, *realMax, *imagMin)
	currentFilename := ""
	i := 1
	for i <= zoom.steps {
		representation := config.representation()
		currentFilename = zoom.name(i, fileNameValue)
		exporter, exporterErr := CreateExporter(*exporterName, representation, *folder, currentFilename, coloring)
		if exporterErr != nil {
			fmt.Print(exporterErr)
			os.Exit(1)
		}
		result, err := exporter.export()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		config = zoom.update(config)
		i += 1

		fmt.Println(result)
	}

}
