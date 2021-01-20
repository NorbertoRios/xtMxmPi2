package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
	"reflect"
)

var HeartBitHeader = [12]byte{0x08, 0x16, 0x02, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x52, 0x00, 0x00, 0x00}

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
	if isBinaryPayload(buffer) {
		fmt.Println("got keep_alive")
	} else {
		responseJ, _ := json.Marshal(createValidHeatBitResponse(h.SESSION))
		bytes := append(validMagicPackageHeader[:], responseJ...)
		err := channel.SendBytes(bytes)
		fmt.Printf("\nsent packet back as text: %s", bytes)
		if err != nil {
			fmt.Errorf("HandleRequest %e", err)
		}
	}
}

func (h HeartBit) ParseDtoFromData(buffer []byte) interface{} {
	if isBinaryPayload(buffer) {
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

func isBinaryPayload(buffer []byte) bool {
	return reflect.DeepEqual(HeartBitHeader[:], buffer[:12])
}
