package main

import (
	"testing"
)

func TestCreateIncorrectStepsZoom(t *testing.T) {
	tests := []struct {
		steps int
	}{
		{0},
		{-1},
	}
	for _, test := range tests {
		zoom, err := CreateZoom(test.steps, 0.0, 0.0, 0.0)
		if err == nil {
			t.Errorf("Expecting error when incorrect step(%d) in CreateZoom; got: nil", test.steps)
		}
		if 1 != zoom.steps {
			t.Errorf("Incorrect steps when creating default Zoom; got: %d, expected: 1", zoom.steps)
		}
	}
}

func TestCreateIncorrectRatioZoom(t *testing.T) {
	tests := []struct {
		ratio float64
	}{
		{1.001},
		{-1.001},
	}
	for _, test := range tests {
		zoom, err := CreateZoom(1, test.ratio, 0.0, 0.0)
		if err == nil {
			t.Errorf("Expecting error when incorrect ratio(%f) in CreateZoom; got: nil", test.ratio)
		}
		if 0 != zoom.ratio {
			t.Errorf("Incorrect ratio when creating default Zoom; got: %f, expected: 0.0", zoom.ratio)
		}
	}
}

func TestCreateNonChangingZoom(t *testing.T) {
	_, err := CreateZoom(2, 0, 0.0, 0.0)
	if err == nil {
		t.Errorf("Expecting error when ratio is 0 and steps > 1; got: nil")
	}
}

func TestZoomGeneratesNewConfig(t *testing.T) {
	tests := []struct {
		steps int
		ratio float64

		expectedRealMin float64
		expectedRealMax float64
		expectedImagMin float64
	}{
		{1, 0.0, -2.0, 2.0, -1.0},
		{2, 0.5, -1.0, 1.0, -0.5},
		{2, 0.1, -1.8, 1.8, -0.9},
	}
	size, _ := CreateSize(100, 50, 1)
	config := CreateConfig(size, 10, -2.0, 2.0, -1.0)
	for i, test := range tests {
		zoom, _ := CreateZoom(test.steps, test.ratio, 0.0, 0.0)
		newConfig := zoom.update(config)
		if newConfig.realMin != test.expectedRealMin {
			t.Errorf("%d. Incorrect realMin in updated config; got: %f, expected: %f", i, newConfig.realMin, test.expectedRealMin)
		}
		if newConfig.realMax != test.expectedRealMax {
			t.Errorf("%d. Incorrect realMax in updated config; got: %f, expected: %f", i, newConfig.realMax, test.expectedRealMax)
		}
		if newConfig.imagMin != test.expectedImagMin {
			t.Errorf("%d. Incorrect realMin in updated config; got: %f, expected: %f", i, newConfig.imagMin, test.expectedImagMin)
		}
	}
}

func TestZoomGeneratesNewConfigWithCalculatedCenter(t *testing.T) {
	tests := []struct {
		realMin float64
		realMax float64
		imagMin float64

		realCenter float64
		imagCenter float64
		ratio      float64

		expectedRealMin float64
		expectedRealMax float64
		expectedImagMin float64
	}{
		{-2.0, 2.0, -1.0, 0.0, 0.0, 0.5, -1.0, 1.0, -0.5},
		{-2.0, 2.0, -1.0, -1.0, 0.5, 0.5, -1.5, 0.5, -0.25},
		{-2.0, 2.0, -1.0, 1.0, 0.5, 0.5, -0.5, 1.5, -0.25},
		{-2.0, 2.0, -1.0, 1.0, -0.5, 0.5, -0.5, 1.5, -0.75},
		{-2.0, 2.0, -1.0, -1.0, -0.5, 0.5, -1.5, 0.5, -0.75},
	}
	size, _ := CreateSize(100, 50, 1)
	for i, test := range tests {
		config := CreateConfig(size, 10, test.realMin, test.realMax, test.imagMin)
		zoom, _ := CreateZoom(2, test.ratio, test.realCenter, test.imagCenter)
		newConfig := zoom.update(config)

		if newConfig.realMin != test.expectedRealMin {
			t.Errorf("%d. Incorrect realMin in updated config; got: %f, expected: %f", i, newConfig.realMin, test.expectedRealMin)
		}
		if newConfig.realMax != test.expectedRealMax {
			t.Errorf("%d. Incorrect realMax in updated config; got: %f, expected: %f", i, newConfig.realMax, test.expectedRealMax)
		}
		if newConfig.imagMin != test.expectedImagMin {
			t.Errorf("%d. Incorrect imagMin in updated config; got: %f, expected: %f", i, newConfig.imagMin, test.expectedImagMin)
		}
	}
}

func TestZoomGenerateFileName(t *testing.T) {
	tests := []struct {
		steps        int
		step         int
		name         string
		expectedName string
	}{
		{1, 1, "", "1"},
		{1, 1, "lorem", "lorem-1"},
		{10, 4, "lorem", "lorem-04"},
		{10, 4, "", "04"},
		{100, 25, "lorem", "lorem-025"},
		{99, 25, "lorem", "lorem-25"},
	}
	name := ""
	for _, test := range tests {
		zoom, err := CreateZoom(test.steps, 0.1, 0.0, 0.0)
		if err != nil {
			t.Errorf("Unexpected error creating zoom with %d steps: %s", test.steps, err)
		}
		name = zoom.name(test.step, test.name)
		if test.expectedName != name {
			t.Errorf("Error creating name in zoom(%d) for name %q; got: %q, expected: %q", test.steps, test.name, name, test.expectedName)
		}
	}
}

func TestZoomGeneratesNewIterations(t *testing.T) {
	tests := []struct {
		steps      int
		ratio      float64
		iterations int
		expected   int
	}{
		{1, 0.0, 50, 50},
		{2, 0.5, 100, 150},
		{2, 0.5, 7, 10},
	}
	size, _ := CreateSize(100, 50, 1)

	for i, test := range tests {
		config := CreateConfig(size, test.iterations, -2.0, 2.0, -1.0)
		zoom, err := CreateZoom(test.steps, test.ratio, 0.0, 0.0)
		if err != nil {
			t.Errorf("%d. Unexpected error creating zoom with %d steps", i, test.steps)
		}
		newConfig := zoom.update(config)
		if newConfig.iterations != test.expected {
			t.Errorf("%d. Error updating config with %d iterations at ratio %f; new iterations got: %d, expected: %d",
				i, test.iterations, test.ratio, newConfig.iterations, test.expected)
		}
	}
}
