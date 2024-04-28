package main

import (
	"fmt"
	"log"
	"strings"
)

func SelectCommand(discord *Discord, event *MessageCreateEvent, text []string) {
	switch strings.ToUpper(text[1]) {
	case "HELP":
		log.Printf("Command help")
		runHelp(discord, event)
	case "HELLDIVE":
		log.Printf("Command Helldive")
		runHelldive(discord, event, text)
	case "SUDOKU":
		log.Printf("Command Sudoku")
	default:
		log.Printf("No command found in mention")
	}
}

func runHelp(discord *Discord, event *MessageCreateEvent) {
	msg := "```\n"
	msg += "HELLDIVE\n"
	msg += "-> PLANETS <planet_name|OPTIONAL>\n"
	msg += "```"

	SendMessage(discord, event.ChannelId, msg)
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
		} else if len(text) > 3 {
			planet := strings.Join(text[3:], " ")
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

func listHelldivers2PlanetInfo(discord *Discord, event *MessageCreateEvent, planet string) {
	info := GetWarCampaignResponse()

	for _, i := range *info {
		if strings.ToUpper(i.Name) == strings.ToUpper(planet) {
			msg := "```"
			msg += fmt.Sprintf("Name: %s\n", i.Name)
			msg += fmt.Sprintf("Faction: %s\n", i.Faction)
			msg += fmt.Sprintf("Players: %d\n", i.Players)
			msg += fmt.Sprintf("Health: %d/%d (%f)\n", i.Health, i.MaxHealth, i.Percentage)
			msg += fmt.Sprintf("Defense: %t\n", i.Defense)
			msg += fmt.Sprintf("Major Order: %t\n", i.MajorOrder)
			msg += fmt.Sprintf("Biome: %s\n", i.Biome.Slug)
			msg += "```"

			SendMessage(discord, event.ChannelId, msg)
			return
		}
	}

	msg := fmt.Sprintf("Planet not found: %s", planet)
	log.Printf("%s", msg)
	SendMessage(discord, event.ChannelId, msg)
}
