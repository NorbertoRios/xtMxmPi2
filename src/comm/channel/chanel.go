package channel

import (
	"time"
)

//IChannel connection client interface
type IChannel interface {
	Send(message string) error
	SendBytes(message []byte) error
	CloseConnection()
	RemoteAddr() string
	RemoteIP() string
	RemotePort() int
	ConnectedAtTs() time.Time
	LastActivity() time.Time
	ReceivedBytes() int64
	TransmittedBytes() int64
}
