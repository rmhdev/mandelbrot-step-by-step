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
