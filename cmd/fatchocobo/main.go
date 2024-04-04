/*
 * Fat Chocobo is discord bot that does random
 * things solely based on things that I find to
 * be intresting at the time.  This bot has no 
 * real purpose outside of that.
 */

package main

import (
	"flag"
	"os"
	"fmt"
)

// Everything loaded in from config file
type Settings struct {
	Token string `json:"token"`
}

// Everything load in from cmdline
type Cmdline struct {
	Config string 
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

// Load in cmdline args from stdin
func setArgs(cmdline *Cmdline) {
	flag.BoolVar(&cmdline.Version, "version", false, "Print current version")
	flag.StringVar(&cmdline.Config, "config", "./config.json", "Config file to use") 

	flag.Parse()
}

func main() {
	cmdline := new(Cmdline)
	setArgs(cmdline)

	if cmdline.Version {
		printVersion()
	}
}
