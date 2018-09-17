package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Exporter interface {
	name() string
	export() (string, error)
}

func CreateExporter(name string, r Representation, folder string, filename string) (Exporter, error) {
	switch name {
	case "text":
		return TextExporter{r}, nil
	case "image":
		return ImageExporter{r, folder, filename}, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid Exporter type '%s'", name))
}

type TextExporter struct {
	representation Representation
}

func (e TextExporter) name() string {
	return "text"
}

func (e TextExporter) export() (string, error) {
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

	return result, nil
}

type ImageExporter struct {
	representation Representation
	folder         string
	filename       string
}

func (e ImageExporter) name() string {
	return "image"
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
	// Check if the filename is correct. If not, fix it.
	filename := strings.TrimSpace(e.filename)
	if "" == filename {
		filename = fmt.Sprintf("%dx%d", e.representation.width(), e.representation.height())
	}
	if !strings.HasSuffix(strings.ToLower(filename), ".png") {
		filename = fmt.Sprintf("%s.png", filename)
	}
	// Create file using folder+filename, and encode de image in it:
	resultFilename := strings.Join([]string{e.folder, filename}, "/")
	f, err := os.OpenFile(resultFilename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	png.Encode(f, image)

	return resultFilename, nil
}
