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
	conn, _, err := websocket.DefaultDialer.Dial(url+"?encoding=json", headers)

	return conn, err
}

func ListenWebsocket(ws *websocket.Conn, output chan<- *GatewayEventPayload) {
	data := new(GatewayEventPayload)

	for {
		if err := ws.ReadJSON(data); err != nil {
			log.Printf("Trouble reading from websocket\n\tError: %s")
		}

		output <- data
	}
}
