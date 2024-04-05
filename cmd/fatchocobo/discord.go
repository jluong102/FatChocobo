package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

import (
	"github.com/gorilla/websocket"
)

const DISCORD_URL string = "https://discord.com/api"

// Gateway Opcodes
const (
	GATEWAY_OPCODE_DISPATCH              = 0
	GATEWAY_OPCODE_HEARTBEAT             = 1
	GATEWAY_OPCODE_IDENTIFY              = 2
	GATEWAY_OPCODE_PRESENCE_UPDATE       = 3
	GATEWAY_OPCODE_VOICE_STATE_UPDATE    = 4
	GATEWAY_OPCODE_RESUME                = 6
	GATEWAY_OPCODE_RECONNECT             = 7
	GATEWAY_OPCODE_REQUEST_GUILD_MEMBERS = 8
	GATEWAY_OPCODE_INVALID_SESSION       = 9
	GATEWAY_OPCODE_HELLO                 = 10
	GATEWAY_OPCODE_HEARTBEAT_ACK         = 11
)

// Gateway Close Event Codes
const (
	GATEWAY_CLOSE_UNKNOWN_ERROR         = 4000
	GATEWAY_CLOSE_UNKNOWN_OPCODE        = 4001
	GATEWAY_CLOSE_DECODE_ERROR          = 4002
	GATEWAY_CLOSE_NOT_AUTHENTICATED     = 4003
	GATEWAY_CLOSE_AUTHENTICATION_FAILED = 4004
	GATEWAY_CLOSE_ALREADY_AUTHENTICATED = 4005
	GATEWAY_CLOSE_INVALID_SEQ           = 4007
	GATEWAY_CLOSE_RATE_LIMITED          = 4008
	GATEWAY_CLOSE_SESSION_TIMED_OUT     = 4009
	GATEWAY_CLOSE_INVALID_SHARD         = 4010
	GATEWAY_CLOSE_SHARDING_REQUIRED     = 4011
	GATEWAY_CLOSE_INVALID_API_VERSION   = 4012
	GATEWAY_CLOSE_INVALID_INTENT        = 4013
	GATEWAY_CLOSE_DISALLOWED_INTENT     = 4014
)

type Discord struct {
	Websocket *websocket.Conn
	token     string
}

// HTTP Responses
type GatewayBotResponse struct {
	Url               string                  `json:"url"`
	Shards            int                     `json:"shards"`
	SessionStartLimit SessionStartLimitObject `json:"session_start_limit"`
}

// Discord JSON sub-objects
type SessionStartLimitObject struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

// Gateway stuff
type GatewayEventPayload struct {
	Op int         `json:"op"` // Gateway Opcode
	D  interface{} `json:"d"`  // Event Data
	S  int         `json:"s"`  // Sequence number
	T  string      `json:"t"`  // Event name
}

// HTTP Requests
func (this Discord) GetGatewayBot() (*http.Response, error) {
	endpoint := DISCORD_URL + "/gateway/bot"

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	return this.MakeRequest(request)
}

// General HTTP request for discord]
func (this Discord) MakeRequest(request *http.Request) (*http.Response, error) {
	request.Header["Authorization"] = []string{fmt.Sprintf("Bot %s", this.token)}
	request.Header["User-Agent"] = []string{"DiscordBot (Fat Chocobo, 0)"}

	client := http.Client{}
	return client.Do(request)
}

// Constructor
func CreateDiscord(token string) *Discord {
	discord := new(Discord)
	discord.token = token

	return discord
}

// Parsing through HTTP respones
func GetDiscordGatewayBot(discord *Discord) (*GatewayBotResponse, error) {
	data := new(GatewayBotResponse)

	response, err := discord.GetGatewayBot()

	if err != nil {
		return nil, err
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status code: %s", response.Status)
	}

	raw, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(raw, data)
	defer response.Body.Close()

	return data, err
}
