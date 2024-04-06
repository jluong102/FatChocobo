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
		log.Printf("Ready Event")
		payload := ParseOpReadyEvent(data.D)
		log.Printf("v -> %d", payload.V)
		log.Printf("User -> %s", payload.User)
		log.Printf("Guilds -> %s", payload.Guilds)
		log.Printf("Session ID -> %s", payload.SessionId)
		log.Printf("Resume Gateway Url -> %s", payload.ResumeGatewayUrl)
		log.Printf("Shard -> %d", payload.Shard)
		log.Printf("Application -> %v", payload.Application)
	default:
		log.Printf("Unknown Opcode: %d", data.Op)
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

func initGatewayHandshake(discord *Discord) {

}
