package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

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

type ImageExporter struct {
	representation Representation
	folder         string
	filename       string
}

func (e ImageExporter) export() (string, error) {
	image := image.NewRGBA(image.Rect(0, 0, e.representation.width(), e.representation.height()))
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	color := black
	for y := 0; y < e.representation.height(); y++ {
		for x := 0; x < e.representation.width(); x++ {
			color = white
			if e.representation.isInside(x, y) {
				color = black
			}
			image.Set(x, y, color)
		}
	}
	// If destination folder does not exist, create it:
	if _, folderErr := os.Stat(e.folder); os.IsNotExist(folderErr) {
		folderErr = os.MkdirAll(e.folder, 0755)
		if folderErr != nil {
			return "", folderErr
		}
	}
	// Create file using folder+filename, and encode de image in it:
	resultFilename := strings.Join([]string{e.folder, e.filename}, "/")
	f, err := os.OpenFile(resultFilename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	png.Encode(f, image)

	return resultFilename, nil
}
