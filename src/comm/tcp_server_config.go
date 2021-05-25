package comm

import (
	"expvar"
	"streamax-go/interfaces"
)

var (
	//ServerCounters composes service metrics
	ServerCounters   = expvar.NewMap("Server")
	DeviceChannelMap map[string]*interfaces.IChannel
)
