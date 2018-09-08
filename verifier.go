package main

type Verifier struct{}

// Returns the next complex number z in "z = z^2 + c".
// Uses the perfect square formula: "(a + b)^2 = a^2 + 2 * a * b + b^2"
func (v Verifier) next(realZ float64, imagZ float64, realC float64, imagC float64) (float64, float64) {
	realNew := realZ*realZ - imagZ*imagZ + realC
	imagNew := 2*realZ*imagZ + imagC

	return realNew, imagNew
}
