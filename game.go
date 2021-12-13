package main

import "log"

type Game struct {
	clients map[*Client]bool
}

func newGame() *Game {
	return &Game{
		make(map[*Client]bool),
	}
}

func (game *Game) canStart() {
	if len(game.clients) == 2 {
		go game.start()
	}
}

func (game *Game) start() {
	log.Println("game started")
	// todo: add game loop
}

func (game *Game) delete(client *Client) {
	delete(game.clients, client)
}

func (game *Game) canAddNewClient() bool {
	return len(game.clients) < 2
}

func (game *Game) add(client *Client) {
	game.clients[client] = true

	game.canStart()
}
