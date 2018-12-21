package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestStdOutput(t *testing.T) {
	representation := NewDefaultVeritication()

	exporter, _ := CreateExporter("text", representation, "", "", DefaultColoring(), DefaultProgress())
	result, _ := exporter.export()
	result = strings.Replace(result, fmt.Sprintln(""), "_", -1)

	expected := "····_****_····_"
	if result != expected {
		t.Errorf("Incorrect export! got: `%s`, expected: `%s`", result, expected)
	}
}

func DefaultColoring() Coloring {
	c, _ := CreateColoring("basic", BlackWhitePalette{})

	return c
}

func DefaultProgress() Progress {
	p, _ := CreateProgress("quiet")

	return p
}

func NewDefaultVeritication() Representation {
	size, _ := CreateSize(4, 3, 1)
	config := CreateConfig(size, 10, -2.0, 2.0, -1.0)
	representation := CreateRepresentation(config)
	verification := Verification{true, 1, 0.0, 0.0}
	for i := 0; i < size.width; i++ {
		representation.set(i, 1, verification)
	}

	return representation
}

func TestImageCreatesFile(t *testing.T) {
	representation := NewDefaultVeritication()

	dir, err := ioutil.TempDir("", "rmhdev-mandelbrot")
	if err != nil {
		log.Fatal(err)
	}
	exporter := ImageExporter{representation, dir, "mandelbrot.png", DefaultColoring(), DefaultProgress()}
	result, exportErr := exporter.export()
	if exportErr != nil {
		t.Errorf("Unexpected error while exporting! Got: '%t', expected nil", exportErr)
	}
	expectedFilename := strings.Join([]string{dir, "mandelbrot.png"}, "/")
	if result != expectedFilename {
		t.Errorf("Incorrect filename exported! Got: '%s', expected '%s'", result, expectedFilename)
	}
	if _, err2 := os.Stat(expectedFilename); os.IsNotExist(err2) {
		t.Errorf("Error creating Mandelbrot image, got: '%s'", err2)
	}
	defer os.RemoveAll(dir)
}

func TestImageCreationCreatesFolderIfDoesNotExist(t *testing.T) {
	representation := NewDefaultVeritication()

	baseDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	expectedDir := strings.Join([]string{baseDir, "new-folder"}, "/")
	exporter := ImageExporter{representation, expectedDir, "mandelbrot.png", DefaultColoring(), DefaultProgress()}
	_, exportErr := exporter.export()
	if exportErr != nil {
		t.Errorf("Unexpected error while creating Mandelbrot image, got: '%s'", exportErr)
	}
	if _, existsErr := os.Stat(expectedDir); os.IsNotExist(existsErr) {
		t.Errorf("Folder has not been created! Expected: '%s'", expectedDir)

		defer os.RemoveAll(expectedDir)
	}
}

func TestImageCreationGeneratesImageNameWhenEmpty(t *testing.T) {
	size, _ := CreateSize(4, 3, 4)
	representation := NewDefaultVerificationWithSize(size)

	dir, err := ioutil.TempDir("", "unknown-image-name")
	if err != nil {
		log.Fatal(err)
	}
	exporter := ImageExporter{representation, dir, "", DefaultColoring(), DefaultProgress()}
	result, exportErr := exporter.export()
	if exportErr != nil {
		t.Errorf("Unexpected error while exporting! Got: '%t', expected nil", exportErr)
	}
	if !strings.HasSuffix(result, ".png") {
		t.Errorf("If image name is empty, app should generate a *.png name. Got: %s", result)
	}
	expectedName := "4x3.png"
	if !strings.HasSuffix(result, expectedName) {
		t.Errorf("If image name is empty, app should generate a name based on the final size. File: %s, Expected name: %s", result, expectedName)
	}
	defer os.RemoveAll(dir)
}

func TestImageCreationGeneratesImageNameWithCorrectExtension(t *testing.T) {
	representation := NewDefaultVeritication()

	dir, err := ioutil.TempDir("", "unknown-image-extension")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name         string
		expectedName string
	}{
		{"mandelbrot.jpeg", "mandelbrot.jpeg.png"},
		{"mandelbrot-2", "mandelbrot-2.png"},
		{"mandelbrot-3.PNG", "mandelbrot-3.PNG"},
		{" mandelbrot-4 ", "mandelbrot-4.png"},
	}
	for _, test := range tests {
		exporter := ImageExporter{representation, dir, test.name, DefaultColoring(), DefaultProgress()}
		result, _ := exporter.export()
		if !strings.HasSuffix(result, test.expectedName) {
			t.Errorf("Incorrect filename extension. Got: '%s', expected filename: '%s'", result, test.expectedName)
		}
	}
	defer os.RemoveAll(dir)
}

func TestImageCreationGeneratesImageWithCorrectSize(t *testing.T) {

	dir, err := ioutil.TempDir("", "image-size-check")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		width          int
		height         int
		factor         int
		expectedWidth  int
		expectedHeight int
	}{
		{4, 3, 1, 4, 3},
		{4, 3, 10, 4, 3},
	}

	for _, test := range tests {
		size, _ := CreateSize(test.width, test.height, test.factor)
		representation := NewDefaultVerificationWithSize(size)
		exporter := ImageExporter{representation, dir, "image.png", DefaultColoring(), DefaultProgress()}
		imagePath, _ := exporter.export()

		file, err := os.Open(imagePath)
		if err != nil {
			t.Errorf("Cannot check image size: %v", err)
		}

		image, _, err := image.DecodeConfig(file)
		if err != nil {
			t.Errorf("Cannot decode image %s: %v", imagePath, err)
		}
		if image.Width != test.expectedWidth {
			t.Errorf("Incorrect image width %s; got: %d, expected: %d", imagePath, image.Width, test.expectedWidth)
		}
		//return image.Width, image.Height
	}
	defer os.RemoveAll(dir)
}

func NewDefaultVerificationWithSize(size Size) Representation {
	config := CreateConfig(size, 10, -2.0, 2.0, -1.0)
	representation := CreateRepresentation(config)
	verification := Verification{true, 1, 0.0, 0.0}
	for i := 0; i < size.width; i++ {
		representation.set(i, 1, verification)
	}

	return representation
}

func TestCreateTextExporter(t *testing.T) {
	representation := NewDefaultVeritication()
	exporter, _ := CreateExporter("text", representation, "", "", DefaultColoring(), DefaultProgress())
	if "text" != exporter.name() {
		t.Errorf("Incorrect exporter created, expected 'text' exporter")
	}
}

func TestCreateImageExporter(t *testing.T) {
	representation := NewDefaultVeritication()

	baseDir, err := ioutil.TempDir("", "exporter-factory")
	if err != nil {
		log.Fatal(err)
	}
	exporter, _ := CreateExporter("image", representation, baseDir, "mandelbrot.png", DefaultColoring(), DefaultProgress())
	if "image" != exporter.name() {
		t.Errorf("Incorrect exporter created, expected 'image' exporter")
	}
}

func TestCreateUnknownExporter(t *testing.T) {
	representation := NewDefaultVeritication()
	exporter, err := CreateExporter("lorem", representation, "", "", DefaultColoring(), DefaultProgress())
	if err == nil {
		t.Errorf("Creating an incorrect 'lorem' exporter should return error, got: nil")
	}
	if exporter != nil {
		t.Errorf("Creating an incorrect 'lorem' exporter should return nil as exporter")
	}
}
