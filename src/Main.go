package main

import (
	"comm"
	"config"
	"service"
)

func main() {
	c := config.OpenDBConnection()
	service.DB = c
	server := comm.NewTCPServer("", config.MainPort)
	serverVideo := comm.NewTCPServer("", config.VideoServerPort)
	go serverVideo.Listen()
	server.Listen()
}
