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

func (r Representation) isInside(x int, y int) bool {
	return r.points[y][x].isInside
}
