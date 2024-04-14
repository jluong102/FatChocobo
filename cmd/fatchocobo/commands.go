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
		if len(text) == 3 {
			listHelldivers2Planets(discord, event)
		} else if len(text) == 4 {
			listHelldivers2PlanetInfo(discord, event, planet)
		}
	default:
		log.Printf("Unknown Helldive command %s", text[2])
	}
}

func listHelldivers2Planets(discord *Discord, event *MessageCreateEvent) {
	planets := GetWarCampaignPlanets()
	var msg string

	for _, i := range planets {
		msg += i + "\n"
	}

	SendMessage(discord, event.ChannelId, msg)
}
