package main

type Representation struct {
	points [][]Verification
	config Config
}

func CreateRepresentation(config Config) Representation {
	points := make([][]Verification, config.size.rawHeight())
	for i := range points {
		points[i] = make([]Verification, config.size.rawWidth())
	}

	return Representation{points, config}
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
