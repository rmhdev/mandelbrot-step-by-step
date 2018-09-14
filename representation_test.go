package main

import "testing"

func TestDefaultIsInside(t *testing.T) {
	representation := CreateRepresentation(11, 11)
	result := representation.isInside(0, 0)
	if false != result {
		t.Errorf("Incorrect default value for isInside(%d, %d), got: %t, expected: %t", 0, 0, false, result)
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		x                int
		y                int
		value            bool
		expectedIsInside bool
	}{
		{0, 0, true, true},
		{5, 5, false, false},
	}
	for _, test := range tests {
		representation := CreateRepresentation(11, 11)
		representation.set(test.x, test.y, test.value)
		if test.value != representation.isInside(test.x, test.y) {
			t.Errorf("Incorrect set(%d, %d, %t), got: %t, expected: %t", test.x, test.y, test.value, test.value, test.expectedIsInside)
		}
	}
}
