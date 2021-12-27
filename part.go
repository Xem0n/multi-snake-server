package main

type Part struct {
	pos *Point
	size *Point
}

func newPart(pos *Point) *Part {
	return &Part{
		pos,
		&Point{
			40,
			40,
		},
	}
}