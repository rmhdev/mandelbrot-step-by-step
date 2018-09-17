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
	representation := CreateRepresentation(4, 3)
	representation.set(0, 1, true)
	representation.set(1, 1, true)
	representation.set(2, 1, true)
	representation.set(3, 1, true)

	exporter := Exporter{representation}
	result := strings.Replace(exporter.export(), fmt.Sprintln(""), "_", -1)

	expected := "····_****_····_"
	if result != expected {
		t.Errorf("Incorrect export! got: `%s`, expected: `%s`", result, expected)
	}
}

func TestImageCreatesFile(t *testing.T) {
	representation := CreateRepresentation(4, 3)
	representation.set(0, 1, true)
	representation.set(1, 1, true)
	representation.set(2, 1, true)
	representation.set(3, 1, true)

	dir, err := ioutil.TempDir("", "rmhdev-mandelbrot")
	if err != nil {
		log.Fatal(err)
	}
	exporter := ImageExporter{representation, dir, "mandelbrot.png"}
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
	representation := CreateRepresentation(4, 3)
	representation.set(0, 1, true)
	representation.set(1, 1, true)
	representation.set(2, 1, true)
	representation.set(3, 1, true)

	baseDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	expectedDir := strings.Join([]string{baseDir, "new-folder"}, "/")
	exporter := ImageExporter{representation, expectedDir, "mandelbrot.png"}
	_, exportErr := exporter.export()
	if exportErr != nil {
		t.Errorf("Unexpected error while creating Mandelbrot image, got: '%s'", exportErr)
	}
	if _, existsErr := os.Stat(expectedDir); os.IsNotExist(existsErr) {
		t.Errorf("Folder has not been created! Expected: '%s'", expectedDir)

		defer os.RemoveAll(expectedDir)
	}
}
