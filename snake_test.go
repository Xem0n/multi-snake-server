package main

import "testing"

func TestMove(t *testing.T) {
	snake := newSnake()

	prevPos := make([]*Point, len(snake.parts))

	for i, part := range snake.parts {
		prevPos[i] = part.pos
	}

	snake.direction = Up
	snake.move()

	for i := 0; i < len(prevPos) - 2; i++ {
		if prevPos[i] != snake.parts[i + 1].pos {
			t.Errorf("one of the parts hasn't moved\nindex = %v\nprevPos = %#v\ncurPos = %#v", i, prevPos[i], snake.parts[i + 1].pos)
			t.FailNow()
		}
	}
}

func TestMoveUp(t *testing.T) {
	snake := newSnake()

	prevPos := snake.parts[0].pos
	snake.direction = Up
	snake.move()

	if prevPos.y <= snake.parts[0].pos.y {
		t.Errorf("invalid move, snake didnt go up\nprevPos = %v\ncurPos = %v", prevPos, snake.parts[0].pos)
	}
}

func TestMoveDown(t *testing.T) {
	snake := newSnake()

	prevPos := snake.parts[0].pos
	snake.direction = Down
	snake.move()

	if prevPos.y >= snake.parts[0].pos.y {
		t.Errorf("invalid move, snake didnt go down\nprevPos = %v\ncurPos = %v", prevPos, snake.parts[0].pos)
	}
}

func TestMoveLeft(t *testing.T) {
	snake := newSnake()

	prevPos := snake.parts[0].pos
	snake.direction = Left
	snake.move()

	if prevPos.x <= snake.parts[0].pos.x {
		t.Errorf("invalid move, snake didnt go left\nprevPos = %v\ncurPos = %v", prevPos, snake.parts[0].pos)
	}
}

func TestMoveRight(t *testing.T) {
	snake := newSnake()

	prevPos := snake.parts[0].pos
	snake.direction = Right
	snake.move()

	if prevPos.x >= snake.parts[0].pos.x {
		t.Errorf("invalid move, snake didnt go right\nprevPos = %v\ncurPos = %v", prevPos, snake.parts[0].pos)
	}
}