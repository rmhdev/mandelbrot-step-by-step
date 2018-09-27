package main

import "testing"

func TestDefaultIsInside(t *testing.T) {
	representation := CreateRepresentation(11, 11)
	result := representation.get(0, 0)
	if false != result.isInside {
		t.Errorf("Incorrect default value for isInside(%d, %d), got: %t, expected: %t", 0, 0, false, result.isInside)
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		x                int
		y                int
		isInside         bool
		expectedIsInside bool
	}{
		{0, 0, true, true},
		{5, 5, false, false},
	}
	for _, test := range tests {
		representation := CreateRepresentation(11, 11)
		representation.set(test.x, test.y, Verification{test.isInside, 1, 0.0, 0.0})
		result := representation.get(test.x, test.y)
		if test.expectedIsInside != result.isInside {
			t.Errorf("Incorrect set(%d, %d, %t), got: %t, expected: %t", test.x, test.y, test.isInside, result.isInside, test.expectedIsInside)
		}
	}
}
