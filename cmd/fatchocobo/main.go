/*
 * Fat Chocobo is discord bot that does random
 * things solely based on things that I find to
 * be intresting at the time.  This bot has no 
 * real purpose outside of that.
 */

package main

import (
	"flag"
)

// Everything loaded in from config file
type Settings struct {
	Token string `json:"token"`
}

// Everything load in from cmdline
type Cmdline struct {
	Config string 
}

// Load in cmdline args from stdin
func setArgs(cmdline *Cmdline) {
	flag.StringVar(&cmdline.Config, "config", "./config.json", "Config file to use") 

	flag.Parse()
}

func main() {
	cmdline := new(Cmdline)
	setArgs(cmdline)
}
