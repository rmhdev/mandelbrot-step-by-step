package main

type Representation struct {
	points [][]Verification
	size   Size
}

func CreateRepresentation(size Size) Representation {
	points := make([][]Verification, size.rawHeight())
	for i := range points {
		points[i] = make([]Verification, size.rawWidth())
	}

	return Representation{points, size}
}

func (r Representation) cols() int {
	return len(r.points[0])
}

func (r Representation) rows() int {
	return len(r.points)
}

func (r Representation) set(x int, y int, v Verification) {
	r.points[y][x] = v
}

func (r Representation) get(x int, y int) Verification {
	return r.points[y][x]
}
