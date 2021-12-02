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

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
