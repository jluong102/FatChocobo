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

		payload := ParseOpHelloEvent(data.D)
		log.Printf("heartbeat -> %d", payload.HeartbeatInterval)
	}
}
