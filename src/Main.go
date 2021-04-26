package main

import (
	"comm"
	"config"
	"log"
	"net/http"
	"service"
)

func main() {
	c := config.OpenDBConnection()
	service.DB = c
	server := comm.NewTCPServer("", config.MainPort)
	serverVideo := comm.NewTCPServer("", config.VideoServerPort)
	go serverVideo.Listen()
	go listenHttp()
	server.Listen()

}

func listenHttp() {
	muxer := config.CreateHttpMuxer()
	err := http.ListenAndServe(":8082", muxer)
	if err != nil {
		log.Fatal(err)
	}
}
