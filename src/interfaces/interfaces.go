package interfaces

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
