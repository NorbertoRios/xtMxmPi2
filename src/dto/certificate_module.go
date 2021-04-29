package dto

import (
	"container/list"
	"encoding/json"
	"fmt"
	"reflect"
	"streamax-go/interfaces"
)

type CertificateModule struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            interface{} `json:",omitempty"`
	SESSION              string
	RESPONSE             interface{}
}

func (d CertificateModule) HandleRequest(channel interfaces.IChannel, buffer []byte) {
	d.switchOnOperation(channel)
}

func (d *CertificateModule) switchOnOperation(channel interfaces.IChannel) {
	switch d.OPERATION {
	case "CREATESTREAM":
		d.handleOperationCreateStream(channel)
	case "CONNECT":
		d.handleOperationConnect(channel)
	}
}

func (d *CertificateModule) handleOperationCreateStream(channel interfaces.IChannel) {
	sm := certificateCreateStreamResponse(channel.GetCurrentSession())
	header := &GeneralPackageHeader{}
	marshal, _ := json.Marshal(sm)
	headerBytes := header.toHeaderBytes(uint(len(marshal)))
	channel.SendBytes(append(headerBytes, marshal...))
}

func (d *CertificateModule) handleOperationConnect(channel interfaces.IChannel) {
	rawParam := d.PARAMETER.(map[string]interface{})
	mar, _ := json.Marshal(rawParam)
	var cp *CertificateParameter
	errM := json.Unmarshal(mar, cp)
	if errM != nil && cp != nil {
		d.checkDeviceIdentity(channel, cp.DSNO)
	}
	d.checkDeviceSession(channel, d.SESSION)
	task := GetFirstByDeviceAndResponseType(channel.GetDevice(), CertificateModule{})
	if task != nil {
		task.ProcessResponse(d)
	}
	response := certificateConnectCreateValidResponse(d.SESSION)
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

func (d CertificateModule) checkDeviceIdentity(channel interfaces.IChannel, deviceId string) {
	if channel.GetDevice() == nil {
		var id interfaces.Device = CameraDevice{Id: deviceId}
		_, loaded := DevicesQHolder.LoadOrStore(id, QueueHolder{Q: list.New()})
		if loaded {
			channel.SetDevice(&id)
			//var q = store.(QueueHolder).Q
			fmt.Printf("existing device identity: %v", deviceId)
		} else {
			fmt.Printf("new device identity: %v", deviceId)
		}
	}
}

func (d CertificateModule) checkDeviceSession(c interfaces.IChannel, session string) {
	if c.GetCurrentSession() == "" {
		if session != "" {
			c.SetCurrentSession(session)
		}
	}
}

func GetFirstByDeviceAndResponseType(device *interfaces.Device, rType interface{}) interfaces.Task {
	if device == nil {
		return nil
	}
	value, ok := DevicesQHolder.Load(*device)
	if ok {
		var l = value.(QueueHolder).Q
		front := l.Front()
		for n := front; n.Value != nil; {
			task := n.Value.(interfaces.Task)
			typeV := task.GetType()
			if IsInstanceOf(typeV, rType) {
				return task
			}
			if n.Next() == nil {
				break
			}
			n = n.Next()
		}
	}
	return nil
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func certificateConnectCreateValidResponse(session string) CertificateModule {
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

func certificateCreateStreamResponse(session string) CertificateModule {
	return CertificateModule{
		MODULE:    "CERTIFICATE",
		OPERATION: "CREATESTREAM",
		SESSION:   session,
		RESPONSE: &OperationCreateStreamResponseError{
			ERRORCODE:  0,
			ERRORCAUSE: "",
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

type OperationCreateStreamParameter struct {
	DEVTYPE    string
	DSNO       string
	STREAMNAME string
	VISION     string
}

type OperationCreateStreamResponseError struct {
	ERRORCODE  int
	ERRORCAUSE string
}
