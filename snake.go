package main

const STARTING_LENGTH = 50
const VELOCITY = 1

type Snake struct {
	parts []*Part
}

func newSnake() *Snake {
	// temporary starting pos
	x := 10.
	y := 500.
	parts := make([]*Part, STARTING_LENGTH)

	for i := 0; i < STARTING_LENGTH; i++ {
		parts[i] = newPart(&Point{
			x,
			y,
		})

		y += VELOCITY
	}

	return &Snake{
		parts,
	}
}

func (snake *Snake) think(delta float64) {
	// for _, part := range snake.parts {
	// magic wth part
	// }
}