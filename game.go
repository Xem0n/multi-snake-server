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
	// append(game.clients, client)
}
