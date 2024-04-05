package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const DISCORD_URL string = "https://discord.com/api"

type Discord struct {
	token string
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
	request.Header[""] = []string{"DiscordBot (Fat Chocobo, 0)"}

	client := http.Client{}
	return client.Do(request)
}

// Constructor
func CreateDiscord(token string) *Discord {
	discord := new(Discord)
	discord.token = token

	return discord
}
