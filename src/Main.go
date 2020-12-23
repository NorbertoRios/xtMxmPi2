package main

import "comm"

func main() {
	server := comm.NewTCPServer("", 8080)
	server.Listen()
}
