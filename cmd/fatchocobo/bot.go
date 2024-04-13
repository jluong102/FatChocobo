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
		log.Printf("Dispatch event")
		handleDispatch(data)
		ParseOpReadyEvent(data.D)
	default:
		log.Printf("Unknown Opcode: %d", data.Op)
	}
}

func handleDispatch(data *GatewayEventPayload) {
	switch data.T {
	case "READY":
		ParseOpReadyEvent(data.D)
	case "MESSAGE_CREATE":
		ParseOpMessageCreateEvent(data.D)
	default:
		log.Printf("Unknown dispatch type %s", data.T)
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
