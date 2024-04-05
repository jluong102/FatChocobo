package main

const DISCORD_URL string = "https://discord.com/api"

type Discord struct {
	token string
}

// Constructor 
func CreateDiscord(token string) *Discord {
	discord := new(Discord)
	discord.token = token

	return discord
}
