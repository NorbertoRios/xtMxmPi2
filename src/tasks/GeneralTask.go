package tasks

import "interfaces"

type GeneralTask struct {
	Device    interfaces.Device
	Observers []interfaces.Observer
	Ready     bool
	RType     interface{}
}

func (t GeneralTask) GetDevice() interfaces.Device {
	return t.Device
}

func (t GeneralTask) GetObservers() []interfaces.Observer {
	return t.Observers
}

func (t GeneralTask) IsReady() bool {
	return t.Ready
}

func (t GeneralTask) GetType() interface{} {
	return t.RType
}
