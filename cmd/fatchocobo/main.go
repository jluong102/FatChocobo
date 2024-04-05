/*
 * Fat Chocobo is discord bot that does random
 * things solely based on things that I find to
 * be intresting at the time.  This bot has no
 * real purpose outside of that.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

// Everything loaded in from config file
type Settings struct {
	Token string `json:"token"`
}

// Everything load in from cmdline
type Cmdline struct {
	Config  string
	Version bool
}

// Set these with the makefile
var VERSION string = "UNKNOWN"
var BUILD_DATE string = "UNKNOWN"

// Print the current version and exit
func printVersion() {
	fmt.Printf("*** Fat Chocobo ***\n")
	fmt.Printf("\tVersion: %s\n", VERSION)
	fmt.Printf("\tBuild Date: %s\n", BUILD_DATE)

	os.Exit(NO_ERROR)
}

func printArt() {
	log.Printf("	　　　／\"'￣フ／)　　　　       、")
	log.Printf("　　,/ ,--、 ￣､__フ　　　　 ／/      ")
	log.Printf("　 ,ヘｌ⌒ﾉ 　　＞　　　　,／　/＿	  ")
	log.Printf("　( ＿l_\"_ニ_　く＿　 ／）　／　／    ")
	log.Printf("　 ゛　,＞　　　　　 フ､　　　　､､＞  ")
	log.Printf("　　 <\"　（　　　　　　フ　_＿＞      ")
	log.Printf("　　　ヽ　 ＼､､　＿フ' ノ             ")
	log.Printf("　　　　＼、＿＿､､,_ノ゛			  ")
	log.Printf("　　　　　 　　 〉ﾆ〉ﾆ〉              ")
	log.Printf("　　　　 　 ,､_/ﾆ/ﾆ/                  ")
	log.Printf("　　　　　∠ｌ∠ｌ､ニ＞                 ")
}

// Load in cmdline args from stdin
func setArgs(cmdline *Cmdline) {
	flag.BoolVar(&cmdline.Version, "version", false, "Print current version")
	flag.StringVar(&cmdline.Config, "config", "./config.json", "Config file to use")

	flag.Parse()
}

func checkArgs(cmdline *Cmdline) {

}

// Read in from config file
func loadSettings(configPath string, settings *Settings) {
	// Confirm valid config file path
	if _, err := os.Stat(configPath); err != nil {
		log.Printf("Unable to find config file\n\tError: %s", err)
		os.Exit(MISSING_CONFIG_FILE)
	}

	content, err := os.ReadFile(configPath)

	if err != nil {
		log.Printf("Failed to read from %s\n\tError: %s", configPath, err)
		os.Exit(FILE_READ_ERROR)
	}

	if err := json.Unmarshal(content, settings); err != nil {
		log.Printf("Unable to parse json from %s\n\tError: %s", configPath, err)
		os.Exit(JSON_PARSE_ERROR)
	}
}

// Make sure we have all the settings we need
func checkSettings(settings *Settings) {
	if settings.Token == "" {
		log.Printf("Missing settings \"token\" in config")
		os.Exit(MISSING_SETTING)
	}
}

func initDiscord(discord *Discord) {
	if gateway, err := GetDiscordGatewayBot(discord); err != nil {
		log.Printf("Failed to get discord gateway\n\tError: %s")
		os.Exit(GATEWAY_ERROR)
	} else {
		if discord.Websocket, err = CreateWebsocketConnection(gateway.Url, nil); err != nil {
			log.Printf("Failed to connect to websocket\n\tError: %s", err)
			os.Exit(WEBSOCKET_CONNECTION_ERROR)
		} else {
			log.Printf("Connected to %s", gateway.Url)
		}
	}
}

func main() {
	settings := new(Settings)
	cmdline := new(Cmdline)
	setArgs(cmdline)

	if cmdline.Version {
		printVersion()
	}

	printArt()
	checkArgs(cmdline)
	loadSettings(cmdline.Config, settings)
	checkSettings(settings)

	discord := CreateDiscord(settings.Token)
	initDiscord(discord)
	ListenWebSocket(discord.Websocket)
}
