package controller

import (
	"comm/channel"
	"fmt"
)

func HandleTCPPacket(c channel.IChannel, buffer []byte) {
	fmt.Printf("New packet: %v", buffer)
}
