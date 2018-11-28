package main

import (
	"errors"
	"fmt"
	"strings"
)

type Zoom struct {
	steps      int
	ratio      float64
	realCenter float64
	imagCenter float64
}

func CreateZoom(step int, ratio float64, realCenter float64, imagCenter float64) (Zoom, error) {
	if step <= 0 {
		return Zoom{1, 0.0, 0.0, 0.0}, errors.New(fmt.Sprintf("Step must be greater than 0; got: %d", step))
	}
	if ratio > 1.0 || ratio < -1.0 {
		return Zoom{1, 0.0, 0.0, 0.0}, errors.New(fmt.Sprintf("Ratio is out of bounds; got: %f, expected: [%f..%f]", ratio, -1.0, 1.0))
	}
	if ratio == 0.0 && step > 1 {
		return Zoom{1, 0.0, 0.0, 0.0}, errors.New(fmt.Sprintf("Ratio should not be zero when zoom has many steps(%d); got ratio: %f", step, ratio))
	}

	return Zoom{step, ratio, realCenter, imagCenter}, nil
}

func (z Zoom) update(c Config) Config {
	actualCenterReal, actualCenterImag := c.center()
	newCenterReal := (actualCenterReal + z.realCenter) / 2.0
	newCenterImag := (actualCenterImag + z.imagCenter) / 2.0

	actualRadiusReal, actualRadiusImag := c.radius()
	newRadiusReal := actualRadiusReal - (actualRadiusReal * z.ratio)
	newRadiusImag := actualRadiusImag - (actualRadiusImag * z.ratio)

	return CreateConfig(c.size,
		c.iterations+int(float64(c.iterations)*z.ratio),
		newCenterReal-newRadiusReal,
		newCenterReal+newRadiusReal,
		newCenterImag-newRadiusImag)
}

func (z Zoom) name(step int, name string) string {
	pattern := []string{"%0", fmt.Sprintf("%d", len(fmt.Sprintf("%d", z.steps))), "d"}
	currentStep := fmt.Sprintf(strings.Join(pattern, ""), step)
	result := ""
	if "" != name {
		result = fmt.Sprintf("%s-", name)
	}
	return fmt.Sprintf("%s%s", result, currentStep)
}
