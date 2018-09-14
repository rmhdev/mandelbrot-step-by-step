package main

import (
	"fmt"
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
