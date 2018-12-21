package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Exporter interface {
	name() string
	export() (string, error)
}

func CreateExporter(name string, r Representation, folder string, filename string, coloring Coloring, progress Progress) (Exporter, error) {
	switch name {
	case "text":
		return TextExporter{r}, nil
	case "image":
		return ImageExporter{r, folder, filename, coloring, progress}, nil
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
	progress       Progress
}

func (e ImageExporter) name() string {
	return "image"
}

func (e ImageExporter) export() (string, error) {
	rect := image.Rect(0, 0, e.representation.cols(), e.representation.rows())
	rawImage := image.NewRGBA(rect)
	progressText := e.filename + ". export"
	total := e.representation.rows()
	previous := 0
	actual := 0
	for y := 0; y < e.representation.rows(); y++ {
		for x := 0; x < e.representation.cols(); x++ {
			rawImage.Set(x, y, e.coloring.color(e.representation.get(x, y)))
		}
		actual = int(e.progress.maxBars * (y + 1) / total)
		if actual > previous {
			e.progress.bar(progressText, e.progress.maxBars, actual)
			previous = actual
		}
	}
	e.progress.write("generating image... ")

	// resize image
	resizedImage := resize.Resize(uint(e.representation.config.size.width), 0, rawImage, resize.Lanczos3)

	// add label to image
	finalImage := image.NewRGBA(resizedImage.Bounds())
	draw.Draw(finalImage, finalImage.Bounds(), resizedImage, image.ZP, draw.Src)
	e.addLabel(finalImage)

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
		filename = fmt.Sprintf("%dx%d", e.representation.config.size.width, e.representation.config.size.height)
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
	png.Encode(f, finalImage) //resizedImage
	e.progress.writeln("ok")

	return resultFilename, nil
}

func (ex ImageExporter) addLabel(exportedImage draw.Image) {
	x := 20
	y := 20
	face := basicfont.Face7x13
	textMargin := 13
	backgroundColor := color.RGBA{255, 255, 255, 255}
	fontColor := color.RGBA{0, 0, 0, 255}
	realCenter, imagCenter := ex.representation.config.center()
	realRadius, imagRadius := ex.representation.config.radius()
	stringCoords := fmt.Sprintf("center: (%e, %e)", realCenter, imagCenter)
	stringRadius := fmt.Sprintf("radius: (%e, %e)", realRadius, imagRadius)
	stringIterations := fmt.Sprintf("max iterations: %d", ex.representation.config.iterations)
	drawer := &font.Drawer{
		Dst:  exportedImage,
		Src:  image.NewUniform(fontColor),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)},
	}
	boxPadding := 2
	// background rectangle
	boundsA, _ := drawer.BoundString(stringCoords)
	boundsB, _ := drawer.BoundString(stringRadius)
	boundsC, _ := drawer.BoundString(stringIterations)
	maxX := math.Max(
		float64(boundsA.Max.X/64),
		math.Max(
			float64(boundsB.Max.X/64),
			float64(boundsC.Max.X/64)))
	rect := image.Rect(
		int(boundsA.Min.X/64)-boxPadding,
		int(boundsA.Min.Y/64)-boxPadding,
		int(maxX)+boxPadding,
		int(boundsC.Max.Y/64)+textMargin*2+boxPadding)
	if rect.Bounds().Max.X > ex.representation.cols() {
		// text is bigger than image length
		return
	}
	if rect.Bounds().Max.Y > ex.representation.rows() {
		// text is bigger than image height
		return
	}
	draw.Draw(exportedImage, rect, &image.Uniform{backgroundColor}, image.ZP, draw.Src)

	// write strings, one per line:
	drawer.DrawString(stringCoords)
	drawer.Dot = fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6((y + textMargin) * 64)}
	drawer.DrawString(stringRadius)
	drawer.Dot = fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6((y + textMargin*2) * 64)}
	drawer.DrawString(stringIterations)
}
