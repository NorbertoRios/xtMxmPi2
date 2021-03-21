package dto

import (
	"encoding/json"
	"fmt"
	"interfaces"
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

		//funcName(channel)

	}
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
