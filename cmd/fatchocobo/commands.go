package main

import (
	"log"
	"strings"
)

func SelectCommand(discord *Discord, event *MessageCreateEvent) {
	text := strings.Split(event.Content, " ")

	switch strings.ToUpper(text[1]) {
	case "HELP":
		log.Printf("Command help")
	default:
		log.Printf("No command found in mention")
	}
}
