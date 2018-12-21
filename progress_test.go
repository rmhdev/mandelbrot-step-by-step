package main

import (
	"bytes"
	"testing"
)

func TestSimpleWrite(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	expected := "Hello world!"
	progress.write(expected)
	got := buffer.String()

	if got != expected {
		t.Errorf("Got %q, expected %q", got, expected)
	}
}

func TestMultipleWrites(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	progress.write("Hello world!")
	progress.write(" How are you?")
	got := buffer.String()
	expected := "Hello world! How are you?"

	if got != expected {
		t.Errorf("Got %q, expected %q", got, expected)
	}
}

func TestWriteln(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	progress.writeln("Hello world!")
	got := buffer.String()
	expected := "Hello world!\n"

	if got != expected {
		t.Errorf("Got %q, expected %q", got, expected)
	}
}

func TestBarEmpty(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	progress.bar("Lorem Ipsum", 100, 0)
	got := buffer.String()
	expected := "\rLorem Ipsum  [>                   ] "

	if got != expected {
		t.Errorf("Incorrect progress bar (0 %%), got %q, expected %q", got, expected)
	}
}

func TestBarHalf(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	progress.bar("Lorem Ipsum", 100, 50)
	got := buffer.String()
	expected := "\rLorem Ipsum  [==========>         ] "

	if got != expected {
		t.Errorf("Incorrect progress bar (50 %%), got %q, expected %q", got, expected)
	}
}

func TestBarFull(t *testing.T) {
	buffer := &bytes.Buffer{}
	progress := Progress{buffer, 20}
	progress.bar("Lorem Ipsum", 100, 100)
	got := buffer.String()
	expected := "\rLorem Ipsum  [====================] "

	if got != expected {
		t.Errorf("Incorrect progress bar (100 %%), got %q, expected %q", got, expected)
	}
}

func TestCreateProgress(t *testing.T) {
	tests := []struct {
		verbosity       string
		isErrorExpected bool
	}{
		{"v", false},
		{"quiet", false},
		{"lorem", true},
	}
	for _, test := range tests {
		_, err := CreateProgress(test.verbosity)
		if test.isErrorExpected {
			if err == nil {
				t.Errorf("Incorrect progress creation: got nil, expected error!")
			}
		}
	}
}
