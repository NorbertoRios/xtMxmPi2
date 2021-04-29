package main

import (
	"comm"
	"config"
	"service"
)

func main() {
	c := config.OpenDBConnection()
	service.DB = c
	server := comm.NewTCPServer("", config.GetConfig().MainPort)
	serverVideo := comm.NewTCPServer("", config.GetConfig().VideoServerPort)
	go serverVideo.Listen()
	go listenHttp()
	server.Listen()

}

func listenHttp() {
	config.CreateGinServer()
}
