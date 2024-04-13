package main

import (
	"log"
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
		handleMessage(discord, event)
	default:
		log.Printf("Unhandled dispatch type %s", data.T)
	}
}

func handleMessage(discord *Discord, event *MessageCreateEvent) {
	if len(event.Mentions) < 1 {
		return
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
