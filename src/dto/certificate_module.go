package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
	"strings"
)

var validMagicPackageHeader = [...]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x93, 0x52, 0x00, 0x00, 0x00}

type CertificateModule struct {
	MODULE    string
	OPERATION string
	PARAMETER CertificateParameter
	SESSION   string
	RESPONSE  CertificateResponseError
}

func (d CertificateModule) HandleRequest(channel channel.IChannel) {
	response := certificateCreateValidResponse(d.SESSION)
	responseJ, _ := json.Marshal(response)
	err := channel.SendBytes(append(validMagicPackageHeader[:], responseJ...))
	if err != nil {
		fmt.Errorf("HandleRequest %e", err)
	}
}

func (d CertificateModule) ParseDtoFromData(buffer []byte) interface{} {
	var wd CertificateModule
	err := json.Unmarshal(buffer[12:strings.LastIndex(string(buffer), "}")], &wd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return wd
}

func certificateCreateValidResponse(session string) CertificateModule {
	return CertificateModule{
		MODULE:    "CERTIFICATE",
		OPERATION: "CONNECT",
		SESSION:   session,
		RESPONSE: CertificateResponseError{
			ERRORCODE:  0,
			ERRORCAUSE: "",
			MASKCMD:    57,
		},
	}
}

type CertificateResponseError struct {
	ERRORCODE  int
	ERRORCAUSE string
	MASKCMD    int
}

type CertificateParameter struct {
	AUTOCAR  string
	AUTONO   string
	CARNUM   string
	CHANNEL  int
	CID      int
	CNAME    string
	CPN      string
	DEVCLASS int
	DEVNAME  string
	DEVTYPE  int
	DSNO     string
	EID      string
	EV       string
	FSV      int
	ICCID    string
	LINENO   string
	MTYPE    int
	NET      int
	PRO      string
	PV       int
	STYPE    int
	TSE      int
	UNAME    string
	UNO      string
}
