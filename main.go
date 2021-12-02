package main

import (
	"flag"
	"log"
	"net/http"
)

var address = flag.String("address", "0.0.0.0", "server address")
var port = flag.String("port", "4002", "server address")

func main() {
	flag.Parse()
	addr := *address + ":" + *port
	log.Println("server running on: " + addr)

	game := newGame()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(game, w, r)
	})

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
