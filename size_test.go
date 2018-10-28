package main

import "testing"

func TestOutOfBoundsSize(t *testing.T) {
	tests := []struct {
		width  int
		height int
	}{
		{0, 10},
		{-1, 10},
		{10, 0},
		{10, -1},
		{10, 5001},
		{5001, 5000},
	}
	for _, test := range tests {
		_, err := CreateSize(test.width, test.height, 1)
		if nil == err {
			t.Errorf(
				"Unexpected result for incorrect size(%d, %d, %d); got nil, expected error",
				test.width, test.height, 1)
		}
	}
}

func TestCorrectSize(t *testing.T) {
	tests := []struct {
		width  int
		height int
		factor int
	}{
		{10, 10, 1},
		{5000, 5000, 16},
	}
	for _, test := range tests {
		size, err := CreateSize(test.width, test.height, test.factor)
		if err != nil {
			t.Errorf("Unexpected error while size creation! Got: %s, expected: nil", err)
		}
		if size.width != test.width {
			t.Errorf("Incorrect width! Got: %d, expected: %d", size.width, test.width)
		}
		if size.height != test.height {
			t.Errorf("Incorrect height! Got: %d, expected: %d", size.height, test.height)
		}
		if size.factor != test.factor {
			t.Errorf("Incorrect detail! Got: %d, expected: %d", size.factor, test.factor)
		}
	}
}

func TestIncorrectDetail(t *testing.T) {
	tests := []struct {
		detail int
	}{
		{0},
		{-1},
		{17},
	}
	for _, test := range tests {
		_, err := CreateSize(10, 10, test.detail)
		if nil == err {
			t.Errorf(
				"Unexpected result for incorrect detail in size(%d, %d, %d); got nil, expected error",
				10, 10, test.detail)
		}
	}
}

func TestRepresentationSizes(t *testing.T) {
	tests := []struct {
		width          int
		height         int
		detail         int
		expectedWidth  int
		expectedHeight int
	}{
		{10, 11, 1, 10, 11},
		{10, 11, 5, 50, 55},
	}
	for _, test := range tests {
		size, err := CreateSize(test.width, test.height, test.detail)
		if nil != err {
			t.Errorf("Unexpected error; got %s, expected nil", err)
		}
		if test.expectedWidth != size.rawWidth() {
			t.Errorf(
				"Unexpected representation width (for width %d, detail %d); Got: %d, expected: %d",
				test.width, test.detail, size.rawWidth(), test.expectedWidth)
		}
	}
}
