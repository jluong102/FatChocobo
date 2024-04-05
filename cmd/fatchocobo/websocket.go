package main

import (
	"log"
	"net/http"
)

import (
	"github.com/gorilla/websocket"
)

func CreateWebsocketConnection(url string, headers http.Header) (*websocket.Conn, error) {
	// We don't care about the http response
	conn, _, err := websocket.DefaultDialer.Dial(url, headers)

	return conn, err
}

func ListenWebSocket(ws *websocket.Conn) {
	data := new(GatewayEventPayload)

	if err := ws.ReadJSON(data); err != nil {
		log.Printf("Trouble reading from websocket\n\tError: %s")
	}

	log.Printf("op -> %d", data.Op)
	log.Printf("d -> %s", data.D)
	log.Printf("s -> %d", data.S)
	log.Printf("t -> %s", data.T)
}
