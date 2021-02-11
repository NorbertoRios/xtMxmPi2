package tasks

import (
	"dto"
)

type OperationQueryFileListRequestTask struct {
	GeneralTask
	Response *dto.QueryFileListResponse
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
