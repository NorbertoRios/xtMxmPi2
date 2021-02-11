package main

import (
	"comm"
	"config"
)

func main() {
	server := comm.NewTCPServer("", config.MainPort)
	serverVideo := comm.NewTCPServer("", config.VideoServerPort)
	go serverVideo.Listen()
	server.Listen()
}
