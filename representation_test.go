package main

import "testing"

func TestDefaultIsInside(t *testing.T) {
	size, _ := CreateSize(11, 11, 1)
	representation := CreateRepresentation(size)
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
	size, _ := CreateSize(11, 11, 1)
	for _, test := range tests {
		representation := CreateRepresentation(size)
		representation.set(test.x, test.y, Verification{test.isInside, 1, 0.0, 0.0})
		result := representation.get(test.x, test.y)
		if test.expectedIsInside != result.isInside {
			t.Errorf("Incorrect set(%d, %d, %t), got: %t, expected: %t", test.x, test.y, test.isInside, result.isInside, test.expectedIsInside)
		}
	}
}

func TestSizeShouldDefendOnFactor(t *testing.T) {
	size, _ := CreateSize(10, 20, 5)
	representation := CreateRepresentation(size)

	if size.rawWidth() != representation.cols() {
		t.Errorf("Incorrect cols in representation; got: %d, expected: %d", representation.cols(), size.rawWidth())
	}
	if size.rawHeight() != representation.rows() {
		t.Errorf("Incorrect rows in representation; got: %d, expected: %d", representation.rows(), size.rawHeight())
	}
}
