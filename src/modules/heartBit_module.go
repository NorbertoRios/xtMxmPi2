package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"streamax-go/interfaces"
)

type HeartBit struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	SESSION              string
	RESPONSE             *HeartBitResponseError `json:",omitempty"`
}

type HeartBitResponseError struct {
	ERRORCODE  int
	ERRORCAUSE string
}

var crunchSingleton bool

func (h HeartBit) HandleRequest(channel interfaces.IChannel, buffer []byte) {
	//if !crunchSingleton {
	//	RequestFile(channel, "58887_video", 3,
	//		"0-0-36", 1, "20210305170005", "20210305170210")
	//	crunchSingleton = true
	//}
	//if !crunchSingleton {
	//	RequestFile(channel, "58895_video", 3,
	//		"0-0-101", 1, "20210401075506", "20210401081417")
	//	crunchSingleton = true
	//}
	OperationQueryFileListRequest(channel,
		"ad54b3ad-1493-41bf-9829-eaf5e1623582",
		"20210401030000",
		"20210401220000", 255)
	//reqeustCapcha(channel)
	//reqeustREQUP(channel)
	//fileUploadStat(channel)

	if IsBinaryHeartBit(buffer) {
		fmt.Println("got keep_alive")
	} else {
		responseJ, _ := json.Marshal(createValidHeatBitResponse(h.SESSION))
		bytes := append(h.toHeaderBytes(uint(len(responseJ))), responseJ...)
		err := channel.SendBytes(bytes)
		fmt.Printf("\nsent packet back as text: %s", bytes)
		if err != nil {
			fmt.Errorf("HandleRequest %e", err)
		}
		fmt.Println("got keep_alive")

		//funcName(channel)

	}
}

func fileUploadOp(channel interfaces.IChannel) {
	request := `{
		 "MODULE": "MEDIASTREAMMODEL",
		 "SESSION": "2a741181-a3b3-422d-9e7b-a2afbab8ec09",
		 "OPERATION": "FILEUPLOAD",
		 "PARAMETER": {
			"DEVID" : ""
			"EVTUUID" : ""
			"CMD" : 0
			"BUSINESSTYPE" : 2
			"SERIAL" : 123
		 }
		}`
	message := []byte(request)
	b := &bytes.Buffer{}
	//write, err := b.Write(message)
	json.Compact(b, message)
	PrintDebugPackageInfo(b.Bytes())
	header := GeneralPackageHeader{}
	headerBytes := header.toHeaderBytes(uint(len(b.Bytes())))
	res := append(headerBytes, b.Bytes()...)
	channel.SendBytes(res)
}

func fileUploadStat(channel interfaces.IChannel) {
	request := `{
  "MODULE": "MEDIASTREAMMODEL",
  "SESSION": "2a741181-a3b3-422d-9e7b-a2afbab8ec09",
  "OPERATION": "DOWNLOADDATA",
  "PARAMETER": {
    "CSRC": "0",
    "SSRC": 65535,
    "STREAMNAME": "mystream1264568",
    "STREAMTYPE": 1,
    "DATATYPE": 2,
    "OFFSETFLAG": 0,
    "OFFSET": 0,
    "STARTT": "20210412050913",
    "ENDT": "20210412061119",
    "NT": 7,
    "IPANDPORT": "192.168.88.253:8081",
    "SERIAL": 46009
  }
}`
	message := []byte(request)
	b := &bytes.Buffer{}
	//write, err := b.Write(message)
	err := json.Compact(b, message)
	print(err)
	PrintDebugPackageInfo(b.Bytes())
	header := GeneralPackageHeader{}
	headerBytes := header.toHeaderBytes(uint(len(b.Bytes())))
	res := append(headerBytes, b.Bytes()...)
	channel.SendBytes(res)
}

func reqeustCapcha(channel interfaces.IChannel) {
	request := `{
  "MODULE": "MEDIASTREAMMODEL",
  "OPERATION": "REQUESTCATCHPIC",
  "PARAMETER": {
    "CHANNEL": 1,
    "CMDTYPE": 3,
    "INTERVAL": 5,
    "IPANDPORT": "192.168.88.253:8081",
    "SEGBEGIN1": "210415070615",
    "SEGEND1": "070818",
    "SEGMENTCOUNT": 1,
    "STREAMNAME": "historysnaptest"
  },
  "SESSION": "00000107-6f72-41ca-9c73-b4717ad05bab"
}`
	message := []byte(request)
	b := &bytes.Buffer{}
	//write, err := b.Write(message)
	err := json.Compact(b, message)
	print(err)
	PrintDebugPackageInfo(b.Bytes())
	header := GeneralPackageHeader{}
	headerBytes := header.toHeaderBytes(uint(len(b.Bytes())))
	res := append(headerBytes, b.Bytes()...)
	channel.SendBytes(res)
}

func reqeustREQUP(channel interfaces.IChannel) {
	request := `{
  "MODULE": "MEDIASTREAMMODEL",
  "OPERATION": "REQUPLOADEVIDENCE",
  "PARAMETER": {
    "TASKID": "0-0-95",
    "CMD": 1,
    "IPANDPORT": "192.168.88.253:8081"
  },
  "SESSION": "00000107-6f72-41ca-9c73-b4717ad05bab"
}`
	message := []byte(request)
	b := &bytes.Buffer{}
	//write, err := b.Write(message)
	err := json.Compact(b, message)
	print(err)
	PrintDebugPackageInfo(b.Bytes())
	header := GeneralPackageHeader{}
	headerBytes := header.toHeaderBytes(uint(len(b.Bytes())))
	res := append(headerBytes, b.Bytes()...)
	channel.SendBytes(res)
}

func PrintDebugPackageInfo(buffer []byte) {
	fmt.Printf("Sending New packet: %X", buffer)
	fmt.Println()
	fmt.Printf("Sending New packet as text: %s", buffer)
	fmt.Println()
}

func (h HeartBit) ParseDtoFromData(buffer []byte) interface{} {
	if IsBinaryHeartBit(buffer) {
		return &HeartBit{
			MODULE:    "CERTIFICATE",
			OPERATION: "KEEPALIVE",
			SESSION:   "",
		}
	} else {
		return h.parseContent(buffer)
	}

}

func (h HeartBit) parseContent(buffer []byte) interface{} {
	var hb HeartBit
	h.FillGeneralPackageHeaderFromPackage(buffer)
	err := json.Unmarshal(h.PayloadBody, &hb)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return hb
}

func createValidHeatBitResponse(session string) HeartBit {
	return HeartBit{
		MODULE:    "CERTIFICATE",
		OPERATION: "KEEPALIVE",
		SESSION:   session,
		RESPONSE: &HeartBitResponseError{
			ERRORCODE:  0,
			ERRORCAUSE: "",
		},
	}
}

func IsBinaryHeartBit(buffer []byte) bool {
	gp := (&GeneralPackageHeader{}).FillGeneralPackageHeaderFromPackage(buffer)
	if gp.PayloadType == 22 && gp.PayloadLen == 124 {
		return true
	} else {
		return false
	}
}

func IsVideo(buffer []byte) bool {
	gp := (&GeneralPackageHeader{}).FillGeneralPackageHeaderFromPackage(buffer)
	if gp.PayloadType == 3 {
		return true
	}
	return false
}
