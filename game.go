package main

import (
	"log"
	"math"
	"time"
)

const MAX_DELTA = 50

type Game struct {
	clients map[*Client]bool
}

func newGame() *Game {
	return &Game{
		make(map[*Client]bool),
	}
}

func (game *Game) start() {
	log.Println("game started")

	game.loop()
}

func (game *Game) loop() {
	defer func() {
		game.onStopped()
	}()

	lastTimestamp := time.Now()

	for {
		select {
		// not sure if that tickrate is proper 
		case <-time.After(16 * time.Millisecond):
			if len(game.clients) < 2 {
				break
			}

			curTimestamp := time.Now()
			delta := math.Min(float64(curTimestamp.Sub(lastTimestamp)), MAX_DELTA)
			lastTimestamp = curTimestamp

			game.think(delta);
		}
	}
}

func (game *Game) think(delta float64) {
	for client := range game.clients {
		if client != nil {
			client.snake.think(delta)
		}
	}
}

func (game *Game) onStopped() {
	log.Println("game stopped (someone left the game)")
}

func (game *Game) canStart() bool {
	return len(game.clients) == 2
}

func (game *Game) delete(client *Client) {
	delete(game.clients, client)
}

func (game *Game) canAddNewClient() bool {
	return len(game.clients) < 2
}

func (game *Game) add(client *Client) {
	game.clients[client] = true

	if game.canStart() {
		go game.start()
	}
}
