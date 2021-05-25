package modules

import (
	"encoding/json"
	"fmt"
	"streamax-go/interfaces"
)

type ConfigModel struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            ConfigModelParameter
	SESSION              string
	RESPONSE             ConfigModuleResponse
}

type ConfigModelParameter struct {
	MDVR MDVRParam
}

type MDVRParam struct {
	MAIN []MAINParam
	DOSD DOSDParam
}

type DOSDParam struct {
	ACCE int
	AE   int
	CHN  []string
	COE1 int
	COE2 int
	DE   int
	GE   int
	NE   int
	OP   []XY
	REN  byte
	SE   byte
	SOEM byte
	TE   byte
	VE   byte
}

type XY struct {
	X int8
	Y int8
}

type MAINParam struct {
	AEN int
	AFR int
	ALT int
	AMT int
	AST int
	BR  int
	BRM int
	ECT int
	EPV int
	FR  int
	FT  int
	KFI int
	LCN int
	QLT int
	RST int
	USE int
	VEN int
}

type ConfigModuleResponse struct {
	ERRORCAUSE string
	ERRORCODE  int
}

func (c ConfigModel) HandleRequest(channel interfaces.IChannel, buffer []byte) {

}

func (c ConfigModel) ParseDtoFromData(buffer []byte) interface{} {
	var cm ConfigModel
	c.FillGeneralPackageHeaderFromPackage(buffer)
	err := json.Unmarshal(c.PayloadBody, &cm)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cm
}
