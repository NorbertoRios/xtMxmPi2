package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
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

func (h HeartBit) HandleRequest(channel channel.IChannel, buffer []byte) {
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
		//OperationQueryFileListRequest(channel,
		//	"ad54b3ad-1493-41bf-9829-eaf5e1623582",
		//	"20210216030000",
		//	"20210216220000", 255)

		//funcName(channel)

	}
}

func funcName(channel channel.IChannel) {
	request := `{
		 "MODULE": "MEDIASTREAMMODEL",
		 "SESSION": "2a741181-a3b3-422d-9e7b-a2afbab8ec09",
		 "OPERATION": "REQUESTDOWNLOADVIDEO",
		 "PARAMETER": {
		   "PT": 3,
		   "SSRC": 128,
		   "STREAMNAME": "mystream126453",
		   "STREAMTYPE": 1,
		   "RECORDID": "0-0-150",
		   "CHANNEL": 0,
		   "STARTTIME": "20210202035818",
		   "ENDTIME": "20210202135921",
		   "OFFSETFLAG": 1,
		   "OFFSET": 0,
		   "IPANDPORT": "192.168.88.253:8081",
		   "SERIAL": 46009
		 }
		}`
	message := []byte(request)
	b := &bytes2.Buffer{}
	//write, err := b.Write(message)
	json.Compact(b, message)
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
