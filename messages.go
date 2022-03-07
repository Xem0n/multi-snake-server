package main

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

type MessageHandler interface {
	handle(client *Client)
}

type DirectionMessage struct {
	Type string
	Direction Direction
}

func (message *DirectionMessage) handle(client *Client) {
	if message.Direction == client.snake.direction {
		return
	}

	if int32(message.Direction)&1 == int32(client.snake.direction)&1 {
		return
	}

	if message.Direction > 3 {
		return
	}

	client.snake.direction = message.Direction
}

var messages = map[string]func() MessageHandler {
	"direction": func() MessageHandler {
		return &DirectionMessage{}
	},
}

func decodeMessage(json map[string]interface{}) (MessageHandler, error) {
	typ, ok := json["type"].(string)

	if !ok {
		return nil, errors.New("improper type of message (string required)")
	}

	messageCaller := messages[typ]

	if messageCaller == nil {
		return nil, errors.New("given type does not exist")
	}

	message := messageCaller()
	err := mapstructure.Decode(json, message)

	if err != nil {
		return nil, errors.New("couldnt decode the message")
	}

	return message, nil
}