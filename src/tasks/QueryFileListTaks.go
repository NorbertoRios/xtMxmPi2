package tasks

import (
	"streamax-go/modules"
)

type OperationQueryFileListRequestTask struct {
	GeneralTask
	Response *modules.QueryFileListResponse
}

func (r *OperationQueryFileListRequestTask) GetResponse() interface{} {
	return r.Response
}

func (r *OperationQueryFileListRequestTask) Process() {
	//do smth
}

func (r *OperationQueryFileListRequestTask) ProcessResponse(response interface{}) {
	r.Ready = true
}
