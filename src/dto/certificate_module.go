package dto

import (
	"comm/channel"
	"controller"
	"encoding/json"
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
)

type CertificateModule struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            *CertificateParameter `json:",omitempty"`
	SESSION              string
	RESPONSE             *CertificateResponseError
}

func (d CertificateModule) HandleRequest(channel channel.IChannel, buffer []byte) {
	response := certificateCreateValidResponse(d.SESSION)
	responseJ, _ := json.Marshal(response)
	bytes := append(d.toHeaderBytes(uint(len(responseJ))), responseJ...)
	err := channel.SendBytes(bytes)
	fmt.Printf("\nsent packet back as text: %s", bytes)
	if err != nil {
		fmt.Errorf("HandleRequest %e", err)
	}
}

func (d CertificateModule) ParseDtoFromData(buffer []byte) interface{} {
	d.FillGeneralPackageHeaderFromPackage(buffer)
	var wd CertificateModule
	err := json.Unmarshal(d.PayloadBody, &wd)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if wd.OPERATION == "KEEPALIVE" {
		return HeartBit{}.ParseDtoFromData(buffer)
	}
	return wd
}

func (d CertificateModule) checkDeviceIdentity(channel channel.IChannel, deviceId string) {
	if channel.GetDevice() == nil {
		store, loaded := controller.DevicesQHolder.LoadOrStore(deviceId, queue.New(64))
		if loaded {
			q := store.(queue.Queue)
			q.Dispose()
		}
	}
}

func certificateCreateValidResponse(session string) CertificateModule {
	return CertificateModule{
		MODULE:    "CERTIFICATE",
		OPERATION: "CONNECT",
		SESSION:   session,
		RESPONSE: &CertificateResponseError{
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
