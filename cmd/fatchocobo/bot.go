package main

import (
	"log"
)

func StartBot(discord *Discord) {
	log.Printf("Starting bot...")
	output := make(chan *GatewayEventPayload)

	go ListenWebsocket(discord.Websocket, output)

	for {
		data := <-output
		log.Printf("op -> %d", data.Op)
		log.Printf("s -> %d", data.S)
		log.Printf("t -> %s", data.T)

		// Handle event based on OPCODE
		switch data.Op {
			case GATEWAY_OPCODE_HELLO:
				log.Printf("Hello event received")
				payload := ParseOpHelloEvent(data.D)

				log.Printf("Heartbeat interval: %d", payload.HeartbeatInterval)
				discord.Heartbeat = payload.HeartbeatInterval
				go discord.SendHeartbeatEndless(data.S)
			case GATEWAY_OPCODE_HEARTBEAT_ACK:
				log.Printf("Heartbeat acknowledged")
			default:
				log.Printf("Unknown Opcode: %d", data.Op)
		}
	}
}

