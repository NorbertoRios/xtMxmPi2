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
	ParseDtoFromData(buffer []byte) ModuleHandler
}

func HandleTCPPacket(c channel.IChannel, buffer []byte) {
	printDebugPackageInfo(buffer)
	moduleName := ParseModuleName(buffer)
	typeHandler := *ModuleMap[moduleName]
	typeHandler.ParseDtoFromData(buffer).HandleRequest(c)
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
		return string(buffer[22:strings.Index(string(buffer[22:40]), "\"")])
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
