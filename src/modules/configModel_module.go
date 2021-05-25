package modules

import (
	"encoding/json"
	"fmt"
	"streamax-go/dto"
	"streamax-go/interfaces"
)

type ConfigModel struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            dto.ConfigModelParameter
	SESSION              string
	RESPONSE             dto.ConfigModuleResponse
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
