package main

import "testing"

func TestMinPixelToReal(t *testing.T) {
	realMin := -2.5
	config := Config{10, 10, realMin, 1.0, -1.0, 1.0}
	result := config.toReal(0)
	if result != realMin {
		t.Errorf(
			"Incorrect config.toReal(%d), got: (%f), expected: (%f)",
			0, result, realMin)
	}
}

func TestLastPixelToReal(t *testing.T) {
	realMax := 1.0
	config := Config{10, 10, -2.5, realMax, -1.0, 1.0}
	result := config.toReal(9)
	if result != realMax {
		t.Errorf(
			"Incorrect config.toReal(%d), got: (%f), expected: (%f)",
			9, result, realMax)
	}
}

func TestSomePixelToReal(t *testing.T) {
	config := Config{11, 11, -2.5, 2.5, -1.0, 1.0}
	result := config.toReal(5)
	expected := 0.0
	if result != expected {
		t.Errorf(
			"Incorrect config.toReal(%d), got: (%f), expected: (%f)",
			4, result, expected)
	}
}
