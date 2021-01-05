package controller

import (
	"comm/channel"
	"fmt"
)

func HandleTCPPacket(c channel.IChannel, buffer []byte) {
	fmt.Printf("New packet: %X", buffer)
	fmt.Println()
	fmt.Printf("New packet as text: %s", buffer)
	fmt.Println()
}
