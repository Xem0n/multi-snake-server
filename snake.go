package main

const STARTING_LENGTH = 50
const VELOCITY = 1

type Snake struct {
	parts     []*Part
	direction Direction
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
		Up,
	}
}

func (snake *Snake) think(delta float64) {
	snake.move()
}

func (snake *Snake) move() {
	for i := len(snake.parts) - 1; i >= 1; i-- {
		*snake.parts[i] = *snake.parts[i - 1]
	}

	snake.parts[0].pos = snake.getNewPos()
}

func (snake *Snake) getNewPos() *Point {
	x := float64(int32(snake.direction)&1) * VELOCITY
	x *= 2 - float64(snake.direction)

	y := float64(int32(snake.direction+1)&1) * VELOCITY
	y *= 2 - float64(snake.direction+1)

	pos := snake.parts[0].pos

	return &Point{
		pos.x + x,
		pos.y + y,
	}
}