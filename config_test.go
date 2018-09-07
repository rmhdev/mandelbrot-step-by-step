package main

import "testing"

func TestToReal(t *testing.T) {
	tests := []struct {
		x        int
		expected float64
	}{
		{0, -2.5},
		{10, 2.5},
		{5, 0.0},
	}
	config := Config{11, 11, -2.5, 2.5, -1.0, 1.0}
	for _, test := range tests {
		result := config.toReal(test.x)
		if result != test.expected {
			t.Errorf(
				"Incorrect config.toReal(%d), got: (%f), expected: (%f)",
				test.x, result, test.expected)
		}
	}
}
