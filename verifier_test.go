package main

import "testing"

func TestNextValue(t *testing.T) {
	tests := []struct {
		realZ        float64
		imagZ        float64
		realC        float64
		imagC        float64
		realExpected float64
		imagExpected float64
	}{
		{0.0, 0.0, 1.0, -1.0, 1.0, -1.0},
		{-0.75, 0.75, -0.75, 0.75, -0.75, -0.375},
		{-0.75, -0.375, -0.75, 0.75, -0.328125, 1.3125},
	}
	verifier := Verifier{}
	for _, test := range tests {
		realResult, imagResult := verifier.next(test.realZ, test.imagZ, test.realC, test.imagC)
		if (test.realExpected != realResult) || (test.imagExpected != imagResult) {
			t.Errorf("Incorrect next number on (%f, %f)^2 + (%f, %f). Got: (%f, %f), expected: (%f, %f)",
				test.realZ, test.imagZ, test.realC, test.imagC, realResult, imagResult, test.realExpected, test.imagExpected)
		}
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		partReal           float64
		partImag           float64
		iterations         int
		expectedIsInside   bool
		expectedIterations int
	}{
		{0.0, 0.0, 1, true, 1},
		{-2.5, 1.0, 2, false, 1},
		{-2.5, 1.0, 1, true, 1},
		{-0.75, 0.75, 4, true, 4},
		{-0.75, 0.75, 5, false, 4},
	}
	for _, test := range tests {
		verifier := Verifier{test.iterations}
		result := verifier.verify(test.partReal, test.partImag)
		if result.isInside != test.expectedIsInside {
			t.Errorf("Incorrect! 'Is inside' verification for (%f, %f), got: %t, expected: %t",
				test.partReal, test.partImag, result.isInside, test.expectedIsInside)
		}
		if result.iterations != test.expectedIterations {
			t.Errorf("Incorrect! 'iterations' verification for (%f, %f), got: %d, expected: %d",
				test.partReal, test.partImag, result.iterations, test.expectedIterations)
		}
	}
}
