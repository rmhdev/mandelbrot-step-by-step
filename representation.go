package main

type Representation struct {
	points [][]Verification
}

func CreateRepresentation(width int, height int) Representation {
	points := make([][]Verification, height)
	for i := range points {
		points[i] = make([]Verification, width)
	}

	return Representation{points}
}

func (r Representation) width() int {
	return len(r.points[0])
}

func (r Representation) height() int {
	return len(r.points)
}

func (r Representation) set(x int, y int, v Verification) {
	r.points[y][x] = v
}

func (r Representation) get(x int, y int) Verification {
	return r.points[y][x]
}
