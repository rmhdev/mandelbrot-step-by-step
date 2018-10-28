package main

import "testing"

func TestToReal(t *testing.T) {
	tests := []struct {
		width           int
		height          int
		factor          int
		x               int
		shouldHaveError bool
		expected        float64
	}{
		{11, 11, 1, 0, false, -2.5},
		{11, 11, 1, 10, false, 2.5},
		{11, 11, 1, 5, false, 0.0},
		{11, 11, 1, -1, true, 0.0},
		{11, 11, 1, 11, true, 0.0},
		{11, 11, 2, 0, false, -2.5},
		{11, 11, 2, 21, false, 2.5},
		{11, 11, 4, 43, false, 2.5},
	}
	for _, test := range tests {
		size, _ := CreateSize(test.width, test.height, test.factor)
		config := Config{size, -2.5, 2.5, -1.0, 1.0}
		result, err := config.toReal(test.x)
		if test.shouldHaveError && err == nil {
			t.Errorf("Incorrect, error expected for config.toReal(%d)", test.x)
		}
		if result != test.expected {
			t.Errorf(
				"Incorrect config.toReal(%d) on %dx%d factor %d, got: (%f), expected: (%f)",
				test.x, test.width, test.height, test.factor, result, test.expected)
		}
	}
}

func TestToImag(t *testing.T) {
	tests := []struct {
		width    int
		height   int
		factor   int
		y        int
		hasError bool
		expected float64
	}{
		{11, 11, 1, 0, false, 1.0},
		{11, 11, 1, 10, false, -1.0},
		{11, 11, 1, 5, false, 0.0},
		{11, 11, 1, -1, true, 0.0},
		{11, 11, 1, 11, true, 0.0},
		{11, 11, 2, 0, false, 1.0},
		{11, 11, 2, 21, false, -1.0},
		{11, 11, 4, 43, false, -1.0},
	}

	for _, test := range tests {
		size, _ := CreateSize(test.width, test.height, test.factor)
		config := Config{size, -2.5, 1.0, -1.0, 1.0}
		result, err := config.toImag(test.y)
		if test.hasError && err == nil {
			t.Errorf("Incorrect, error expected for config.toImag(%d)", test.y)
		}
		if result != test.expected {
			t.Errorf(
				"Incorrect config.toImag(%d), got: (%f), expected: (%f)",
				test.y, result, test.expected)
		}
	}
}

func TestRepresentation(t *testing.T) {
	size, _ := CreateSize(3, 3, 1)
	config := Config{size, -2.5, 1.0, -1.0, 1.0}
	verifier := Verifier{1}
	verification := Verification{true, 1, 0.0, 0.0}
	size, _ = CreateSize(11, 11, 1)
	expected := CreateRepresentation(size)
	expected.set(0, 0, verification)
	expected.set(0, 1, verification)
	expected.set(0, 2, verification)
	expected.set(1, 0, verification)
	expected.set(1, 1, verification)
	expected.set(1, 2, verification)
	expected.set(2, 0, verification)
	expected.set(2, 1, verification)
	expected.set(2, 2, verification)

	result := config.representation(verifier)
	for x := 0; x < result.cols(); x++ {
		for y := 0; y < result.rows(); y++ {
			if expected.get(x, y).isInside != result.get(x, y).isInside {
				t.Errorf("Incorrect representation for point(%d, %d), got: %t, expected: %t", x, y, result.get(x, y).isInside, expected.get(x, y).isInside)
			}
		}
	}
}

func TestCreateConfig(t *testing.T) {
	tests := []struct {
		width           int
		height          int
		realMin         float64
		realMax         float64
		imagMin         float64
		expectedImagMax float64
	}{
		{100, 100, -2.0, 2.0, -2.0, 2.0},
		{400, 300, -2.0, 2.0, -2.0, 1.0},
		{1000, 500, -2.0, 2.0, -2.0, 0.0},
	}
	for _, test := range tests {
		size, _ := CreateSize(test.width, test.height, 1)
		config := CreateConfig(size, test.realMin, test.realMax, test.imagMin)
		if test.expectedImagMax != config.imagMax {
			t.Errorf("Incorrect imagMax value. Got: %f, expected: %f", config.imagMax, test.expectedImagMax)
		}
	}
}
