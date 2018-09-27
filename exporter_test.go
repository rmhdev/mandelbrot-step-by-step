package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestStdOutput(t *testing.T) {
	representation := NewDefaultVeritication()

	exporter, _ := CreateExporter("text", representation, "", "", DefaultColoring())
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

func NewDefaultVeritication() Representation {
	representation := CreateRepresentation(4, 3)
	verification := Verification{true, 1, 0.0, 0.0}
	representation.set(0, 1, verification)
	representation.set(1, 1, verification)
	representation.set(2, 1, verification)
	representation.set(3, 1, verification)

	return representation
}

func TestImageCreatesFile(t *testing.T) {
	representation := NewDefaultVeritication()

	dir, err := ioutil.TempDir("", "rmhdev-mandelbrot")
	if err != nil {
		log.Fatal(err)
	}
	exporter := ImageExporter{representation, dir, "mandelbrot.png", DefaultColoring()}
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
	exporter := ImageExporter{representation, expectedDir, "mandelbrot.png", DefaultColoring()}
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
	representation := NewDefaultVeritication()

	dir, err := ioutil.TempDir("", "unknown-image-name")
	if err != nil {
		log.Fatal(err)
	}
	exporter := ImageExporter{representation, dir, "", DefaultColoring()}
	result, exportErr := exporter.export()
	if exportErr != nil {
		t.Errorf("Unexpected error while exporting! Got: '%t', expected nil", exportErr)
	}
	if !strings.HasSuffix(result, ".png") {
		t.Errorf("If image name is empty, app should generate a *.png name. Got: %s", result)
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
		exporter := ImageExporter{representation, dir, test.name, DefaultColoring()}
		result, _ := exporter.export()
		if !strings.HasSuffix(result, test.expectedName) {
			t.Errorf("Incorrect filename extension. Got: '%s', expected filename: '%s'", result, test.expectedName)
		}
	}
	defer os.RemoveAll(dir)
}

func TestCreateTextExporter(t *testing.T) {
	representation := NewDefaultVeritication()
	exporter, _ := CreateExporter("text", representation, "", "", DefaultColoring())
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
	exporter, _ := CreateExporter("image", representation, baseDir, "mandelbrot.png", DefaultColoring())
	if "image" != exporter.name() {
		t.Errorf("Incorrect exporter created, expected 'image' exporter")
	}
}

func TestCreateUnknownExporter(t *testing.T) {
	representation := NewDefaultVeritication()
	exporter, err := CreateExporter("lorem", representation, "", "", DefaultColoring())
	if err == nil {
		t.Errorf("Creating an incorrect 'lorem' exporter should return error, got: nil")
	}
	if exporter != nil {
		t.Errorf("Creating an incorrect 'lorem' exporter should return nil as exporter")
	}
}
