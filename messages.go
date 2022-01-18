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
	// handle message
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