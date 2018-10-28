package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

type Exporter interface {
	name() string
	export() (string, error)
}

func CreateExporter(name string, r Representation, folder string, filename string, coloring Coloring) (Exporter, error) {
	switch name {
	case "text":
		return TextExporter{r}, nil
	case "image":
		return ImageExporter{r, folder, filename, coloring}, nil
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
	for y := 0; y < e.representation.rows(); y++ {
		line := ""
		for x := 0; x < e.representation.cols(); x++ {
			if e.representation.get(x, y).isInside {
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
	coloring       Coloring
}

func (e ImageExporter) name() string {
	return "image"
}

func (e ImageExporter) export() (string, error) {
	rect := image.Rect(0, 0, e.representation.cols(), e.representation.rows())
	rawImage := image.NewRGBA(rect)
	for y := 0; y < e.representation.rows(); y++ {
		for x := 0; x < e.representation.cols(); x++ {
			rawImage.Set(x, y, e.coloring.color(e.representation.get(x, y)))
		}
	}
	// resize image
	resizedImage := resize.Resize(uint(e.representation.size.width), 0, rawImage, resize.Lanczos3)
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
		filename = fmt.Sprintf("%dx%d", e.representation.size.width, e.representation.size.height)
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
	png.Encode(f, resizedImage)

	return resultFilename, nil
}
