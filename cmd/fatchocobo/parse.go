package main

import (
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
	output, ok := data.(HelloEvent)

	if !ok {
		log.Printf("Failed to convert to \"HelloEvent\"")	
	}

	return &output
}
