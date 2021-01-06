package controller

import (
	"comm/channel"
	"dto"
	"fmt"
	"strings"
)

var ModuleMap map[string]*ModuleHandler

type ModuleHandler interface {
	HandleRequest(c channel.IChannel)
	ParseDtoFromData(buffer []byte) interface{}
}

func HandleTCPPacket(c channel.IChannel, buffer []byte) {
	printDebugPackageInfo(buffer)
	moduleName := ParseModuleName(buffer)
	typeHandler := *ModuleMap[moduleName]
	typeHandler.ParseDtoFromData(buffer).(ModuleHandler).HandleRequest(c)
}

func printDebugPackageInfo(buffer []byte) {
	fmt.Printf("New packet: %X", buffer)
	fmt.Println()
	fmt.Printf("New packet as text: %s", buffer)
	fmt.Println()
}

func ParseModuleName(buffer []byte) string {
	mod := string(buffer[14:20])
	if mod == "MODULE" {
		return string(buffer[23 : strings.IndexByte(string(buffer[23:40]), 0x22)+23])
	} else {
		return ""
	}
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
	}
}
