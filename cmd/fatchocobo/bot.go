package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func StartBot(discord *Discord) {
	log.Printf("Starting bot...")
	output := make(chan *GatewayEventPayload)

	go ListenWebsocket(discord.Websocket, output)

	for {
		handleGatewayEvent(discord, <-output)
	}
}

func SendMessage(discord *Discord, channelId Snowflake, msg string) {
	payload := new(CreateMessagePayload)
	payload.Content = msg

	response, err := discord.CreateMessage(channelId, payload)

	if response.StatusCode != 200 {
		log.Printf("Failed to create message: %s", response.Status)
	} else if err != nil {
		log.Printf("Failed to create HTTP request: %s", err)
	} else {
		log.Printf("Message sent to %s", channelId)
	}
}

func handleGatewayEvent(discord *Discord, data *GatewayEventPayload) {
	sendingHeartbeats := false

	// Handle event based on OPCODE
	switch data.Op {
	case GATEWAY_OPCODE_HELLO:
		log.Printf("Hello event received")
		payload := ParseOpHelloEvent(data.D)

		log.Printf("Heartbeat interval: %d", payload.HeartbeatInterval)
		discord.Heartbeat = payload.HeartbeatInterval

		if !sendingHeartbeats {
			go sendEndlessHeartbeats(discord, data.S)
			discord.InitGatewayHandshake(GATEWAY_INTENT_GUILD_MESSAGES)
		} else {
			log.Printf("Already sending heartbeats")
		}

		sendingHeartbeats = true
	case GATEWAY_OPCODE_HEARTBEAT_ACK:
		log.Printf("Heartbeat acknowledged")
	case GATEWAY_OPCODE_DISPATCH:
		log.Printf("Dispatch event received")
		handleDispatch(discord, data)
	default:
		log.Printf("Unknown Opcode: %d", data.Op)
	}
}

func handleDispatch(discord *Discord, data *GatewayEventPayload) {
	log.Printf("Dispatch type %s", data.T)

	switch data.T {
	case "READY":
		event := ParseOpReadyEvent(data.D)
		setBotInfo(discord, event)
	case "MESSAGE_CREATE":
		event := ParseOpMessageCreateEvent(data.D)
		handleMessageCreate(discord, event)
	default:
		log.Printf("Unhandled dispatch type %s", data.T)
	}
}

func handleMessageCreate(discord *Discord, event *MessageCreateEvent) {
	if isMentioned(discord, event) {
		log.Printf("Mention found in %s", event.GuildId)
		handleMention(discord, event)
	}
}

func handleMention(discord *Discord, event *MessageCreateEvent) {
	text := strings.Split(event.Content, " ")

	if text[0] == fmt.Sprintf("<@%s>", discord.BotId) {
		if len(text) == 1 {
			log.Printf("Kweh")
		} else {
			SelectCommand(discord, event, text)
		}
	} else {
		SendMessage(discord, event.ChannelId, "Kweh!")
	}
}

// Auto send heartbeats back before it expires.
// This should be run as a goroutine only.
func sendEndlessHeartbeats(discord *Discord, seq int) {
	for {
		log.Printf("Sending heartbeat")
		discord.SendHeartbeat(seq)
		time.Sleep(time.Duration(discord.Heartbeat)*time.Millisecond - 100)
	}
}

func setBotInfo(discord *Discord, event *ReadyEvent) {
	discord.BotId = event.User.Id
	discord.Username = event.User.Username
}

func isMentioned(discord *Discord, event *MessageCreateEvent) bool {
	for _, i := range event.Mentions {
		if i.Id == discord.BotId {
			return true
		}
	}

	return false
}
