package main

import (
	"log"
	"time"
)

func StartBot(discord *Discord) {
	log.Printf("Starting bot...")
	output := make(chan *GatewayEventPayload)
	sendingHeartbeats := false

	go ListenWebsocket(discord.Websocket, output)

	for {
		data := <-output

		// Handle event based on OPCODE
		switch data.Op {
		case GATEWAY_OPCODE_HELLO:
			log.Printf("Hello event received")
			payload := ParseOpHelloEvent(data.D)

			log.Printf("Heartbeat interval: %d", payload.HeartbeatInterval)
			discord.Heartbeat = payload.HeartbeatInterval

			if !sendingHeartbeats {
				go sendEndlessHeartbeats(discord, data.S)
			} else {
				log.Printf("Already sending heartbeats")
			}

			sendingHeartbeats = true
		case GATEWAY_OPCODE_HEARTBEAT_ACK:
			log.Printf("Heartbeat acknowledged")
		default:
			log.Printf("Unknown Opcode: %d", data.Op)
		}
	}
}

func sendEndlessHeartbeats(discord *Discord, seq int) {
	for {
		log.Printf("Sending heartbeat")
		discord.SendHeartbeat(seq)
		time.Sleep(time.Duration(discord.Heartbeat)*time.Millisecond - 100)
	}
}
