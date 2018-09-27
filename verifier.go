package main

import "math"

type Verifier struct {
	maxIterations int
}

type Verification struct {
	isInside   bool
	iterations int
	realZ      float64
	imagZ      float64
}

func (v Verifier) verify(realC float64, imagC float64) Verification {
	realZ, imagZ, modulusZ := 0.0, 0.0, 0.0
	for i := 0; i < v.maxIterations; i++ {
		modulusZ = math.Sqrt(realZ*realZ + imagZ*imagZ)
		if modulusZ > 2 {
			return Verification{false, i, realZ, imagZ}
		}
		realZ, imagZ = v.next(realZ, imagZ, realC, imagC)
	}
	return Verification{true, v.maxIterations, realZ, imagZ}
}

// Returns the next complex number z in "z = z^2 + c".
// Uses the perfect square formula: "(a + b)^2 = a^2 + 2 * a * b + b^2"
func (v Verifier) next(realZ float64, imagZ float64, realC float64, imagC float64) (float64, float64) {
	realNew := realZ*realZ - imagZ*imagZ + realC
	imagNew := 2*realZ*imagZ + imagC

	return realNew, imagNew
}
