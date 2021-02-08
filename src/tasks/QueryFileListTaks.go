package tasks

import (
	"dto"
	"interfaces"
)

type OperationQueryFileListRequestTask struct {
	Response  *dto.QueryFileListResponse
	Device    interfaces.Device
	Observers []interfaces.Observer
	Ready     bool
}

func (r *OperationQueryFileListRequestTask) GetResponse() interface{} {
	return r.Response
}

func (r *OperationQueryFileListRequestTask) GetDevice() interfaces.Device {
	return r.Device
}

func (r *OperationQueryFileListRequestTask) GetObservers() []interfaces.Observer {
	return r.Observers
}

func (r *OperationQueryFileListRequestTask) Process() {
	//do smth
}

func (r *OperationQueryFileListRequestTask) IsReady() bool {
	return r.Ready
}
