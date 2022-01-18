package main

import "testing"

func TestDecodeMessage(t *testing.T) {
	var json map[string]interface{}
	var err error

	json = map[string]interface{} {
		// "type": "direction",
		"direction": 5,
	}

	_, err = decodeMessage(json)

	if err == nil {
		t.Error("decoded improper message (without type)")
	}

	json = map[string]interface{} {
		"type": 112,
		"direction": 5,
	}

	_, err = decodeMessage(json)

	if err == nil {
		t.Error("decoded improper message (with not string type)")
	}

	json = map[string]interface{} {
		"type": "not existing type",
		"direction": 5,
	}

	_, err = decodeMessage(json)

	if err == nil {
		t.Error("decoded improper message (with not existing type)")
	}

	json = map[string]interface{} {
		"type": "direction",
		"direction": 5,
	}

	_, err = decodeMessage(json)

	if err != nil {
		t.Error("couldn't decode valid json")
	}
}