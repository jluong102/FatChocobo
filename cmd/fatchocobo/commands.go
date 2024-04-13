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
		runHelp(discord, event)
	case "HELLDIVE":
		log.Printf("Command Helldive")
		runHelldive(discord, event, text)
	default:
		log.Printf("No command found in mention")
	}
}

func runHelp(discord *Discord, event *MessageCreateEvent) {

}

func runHelldive(discord *Discord, event *MessageCreateEvent, text []string) {
	if len(text) < 3 {
		SendMessage(discord, event.ChannelId, "Democracy")
		return
	} 

	switch strings.ToUpper(text[2]) {
	case "PLANETS":

	default:
		log.Printf("Unknown Helldive command %s", text[2])
	}	
}
