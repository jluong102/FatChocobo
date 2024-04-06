package main

import (
	"encoding/json"
	"log"
)

/*
 * All "data" interface parameters should be from "d" value
 * that is returned from the discord gateway.
 * The type that should be parsed can be gathered
 * based off of the op code received from the socket.
 * Due to the "d" type being dynamic these parse functions
 * can be used to simplify the process.  Alternativly, you
 * can convert directly from the interface.
 */

func ParseOpHelloEvent(data interface{}) *HelloEvent {
	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Failed to encode data to JSON\n\t%s", err)
		return nil
	}

	output := new(HelloEvent)
	err = json.Unmarshal(encoded, output)

	if err != nil {
		log.Printf("Failed to parse JSON\n\t%s", err)
		return nil
	}

	return output
}

func ParseOpReadyEvent(data interface{}) *ReadyEvent {
	encoded, err := json.Marshal(data)

	if err != nil {
		log.Printf("Failed to encode data to JSON\n\t%s", err)
	}

	output := new(ReadyEvent)
	err = json.Unmarshal(encoded, output)

	if err != nil {
		log.Printf("Failed to parse JSON\n\t%s", err)
		return nil
	}
	return output
}
