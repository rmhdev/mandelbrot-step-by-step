package main

type Representation struct {
	points [][]bool
}

func CreateRepresentation(width int, height int) Representation {
	points := make([][]bool, height)
	for i := range points {
		points[i] = make([]bool, width)
	}

	return Representation{points}
}

func (r Representation) width() int {
	return len(r.points[0])
}

func (r Representation) height() int {
	return len(r.points)
}

func (r Representation) set(x int, y int, isInside bool) {
	r.points[y][x] = isInside
}

func (r Representation) isInside(x int, y int) bool {
	return r.points[y][x]
}
