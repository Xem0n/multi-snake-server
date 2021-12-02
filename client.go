package main

import "net/http"

type Client struct {
	game *Game
}

func serveWebsocket(game *Game, w http.ResponseWriter, r *http.Request) {

}
