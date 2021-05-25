package interfaces

import "time"

type Device interface {
	GetId() string
}

type Observer interface {
	Notify()
}

type Task interface {
	GetResponse() interface{}
	GetDevice() Device
	GetObservers() []Observer
	Process()
	IsReady() bool
	GetType() interface{}
	ProcessResponse(interface{})
}

type AbstractHeader interface {
	FillHeaderFromPackage(buffer []byte) AbstractHeader
	GetPayloadLen() uint
	IsSegmented(buffer []byte) bool
	ContainsAdditionalTCPSegment(buffer []byte) (bool, []byte, []byte)
}

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
	GetPageBuffer() *interface{}
	SetPageBuffer(*interface{})
	GetDSNO() string
	SetDSNO(string)
	GetCurrentSession() string
	SetCurrentSession(s string)
	SetVideoHandler(vh interface{})
	GetVideoHandler() interface{}
}
