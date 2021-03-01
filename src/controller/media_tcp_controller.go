package controller

import (
	"comm/channel"
	"dto"
	"fmt"
	"strings"
	"sync"
)

var ModuleMap map[string]*ModuleHandler

type ModuleHandler interface {
	HandleRequest(c channel.IChannel, buffer []byte)
	ParseDtoFromData(buffer []byte) interface{}
}

func HandleTCPPacket(c channel.IChannel, buffer []byte) {
	//justPrint(buffer)
	moduleName := ParseModuleName(buffer)
	if moduleName == "" {
		moduleName = ParseHeaderBitMaskName(buffer)
		if moduleName == "" {
			return
		}
	}
	if ModuleMap[moduleName] == nil {
		return
	}
	typeHandler := *ModuleMap[moduleName]
	typeHandler.ParseDtoFromData(buffer).(ModuleHandler).HandleRequest(c, buffer)
}

func ParseModuleName(buffer []byte) string {
	mod := string(buffer[14:20])
	if mod == "MODULE" {
		return string(buffer[23 : strings.IndexByte(string(buffer[23:40]), 0x22)+23])
	} else {
		return ""
	}
}

func ParseHeaderBitMaskName(buffer []byte) string {
	if dto.IsBinaryHeartBit(buffer) {
		return "HEARTBIT"
	}
	if dto.IsVideo(buffer) {
		return "VideoHandler"
	}
	return ""
}

func justPrint(tb []byte) {
	fmt.Printf("New packet: %X", tb)
	fmt.Println()
	fmt.Printf("New packet as text: %s", tb)
	fmt.Println()
}

func InitModuleMap() {
	if ModuleMap == nil {
		ModuleMap = make(map[string]*ModuleHandler)
		var mh ModuleHandler = &dto.CertificateModule{
			MODULE: "CERTIFICATE",
		}
		ModuleMap["CERTIFICATE"] = &mh
		var cmmh ModuleHandler = &dto.ConfigModel{
			MODULE: "CONFIGMODEL",
		}
		ModuleMap["CONFIGMODEL"] = &cmmh

		var hbm ModuleHandler = &dto.HeartBit{
			MODULE: "HEARTBIT",
		}
		ModuleMap["HEARTBIT"] = &hbm
		var evm ModuleHandler = &dto.Evem{
			MODULE: "EVEM",
		}
		ModuleMap["EVEM"] = &evm
		var storm ModuleHandler = &dto.Storm{
			MODULE: "STORM",
		}
		ModuleMap["STORM"] = &storm
		var vh ModuleHandler = &dto.VideoHandlerModule{}
		ModuleMap["VideoHandler"] = &vh
	}
	dto.DevicesQHolder = new(sync.Map)
}
