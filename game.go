package main

type Game struct {
	clients []*Client
}

func newGame() *Game {
	return &Game{
		make([]*Client, 0),
	}
}

func (game *Game) add(client *Client) {
	game.clients = append(game.clients, client)
}
