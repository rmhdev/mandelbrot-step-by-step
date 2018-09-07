package main

import "testing"

func TestToReal(t *testing.T) {
	tests := []struct {
		x               int
		shouldHaveError bool
		expected        float64
	}{
		{0, false, -2.5},
		{10, false, 2.5},
		{5, false, 0.0},
		{-1, true, 0.0},
		{11, true, 0.0},
	}
	config := Config{11, 11, -2.5, 2.5, -1.0, 1.0}
	for _, test := range tests {
		result, err := config.toReal(test.x)
		if test.shouldHaveError && err == nil {
			t.Errorf("Incorrect, error expected for config.toReal(%d)", test.x)
		}
		if result != test.expected {
			t.Errorf(
				"Incorrect config.toReal(%d), got: (%f), expected: (%f)",
				test.x, result, test.expected)
		}
	}
}
