package main

import (
	"streamax-go/comm"
	"streamax-go/config"
	"streamax-go/httpService"
	"streamax-go/repository"
	"streamax-go/tcpService"
)

func main() {
	c := config.OpenDBConnection()
	httpService.DB = c
	tcpService.DB = c
	repository.DB = c
	server := comm.NewTCPServer("", config.GetConfig().MainPort)
	serverVideo := comm.NewTCPServer("", config.GetConfig().VideoServerPort)
	go serverVideo.Listen()
	go listenHttp()
	server.Listen()

}

func listenHttp() {
	gs := &comm.GinServer{}
	gs.RunGinServer()
}
