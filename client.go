package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	readWait = time.Millisecond * 500
	writeWait = time.Millisecond * 100
	pingPeriod = time.Millisecond * 100
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	game *Game
	snake *Snake
	closed bool
}

func (client *Client) readPump() {
	defer func() {
		client.close()
	}()

	client.conn.SetReadDeadline(time.Now().Add(readWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(readWait)); return nil })

	for {
		var json map[string]interface{}
		err := client.conn.ReadJSON(&json)

		if err != nil {
			log.Println(err)

			break
		}

		client.handleMessage(json)
	}
}

func (client *Client) handleMessage(json interface{}) {
}

func (client *Client) writePump() {
	pingTicker := time.NewTicker(pingPeriod)

	defer func() {
		pingTicker.Stop()
		client.close()
	}()

	for {
		select {
		// todo: add and handle write channel 
		case <-pingTicker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))

			err := client.conn.WriteMessage(websocket.PingMessage, nil)

			if err != nil {
				return
			}
		}
	}
}

func (client *Client) close() {
	if client.closed {
		return
	}

	client.conn.Close()
	client.game.delete((client))
	client.closed = true

	log.Println("connection closed")
}

func serveWebsocket(game *Game, w http.ResponseWriter, r *http.Request) {
	if (!game.canAddNewClient()) {
		log.Println("game's already started. rejected connection")

		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	client := &Client{
		game: game,
		conn: conn,
		snake: newSnake(),
		closed: false,
	}

	log.Printf("new connection: %s\n", conn.LocalAddr().String())

	go game.add(client)
	go client.readPump()
	go client.writePump()
}
