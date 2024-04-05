package main

import (
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
	ws.ReadJSON()
}
