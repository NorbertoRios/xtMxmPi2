package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
	"strings"
)

type Evem struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            *EvemParameter `json:",omitempty"`
	SESSION              string
	RESPONSE             *EvemResponse `json:",omitempty"`
}

type EvemParameter struct {
	ALARMAS     int
	ALARMCOUNT  int
	ALARMTYPE   int
	ALARMUID    int
	CMDNO       int
	CMDTYPE     int
	CURRENTTIME int
	EVTUUID     string
	L           int
	LEV         int
	P           *struct {
		C int
		J string
		S int
		T string
		V int
		W string
	}
	REAL        int
	RUN         int
	SECNO       int
	SP          int
	ST          int
	TRIGGERTYPE int
}

type EvemResponse struct {
	SERIAL     int
	ERRORCAUSE string
	ERRORCODE  int
	ALARMUID   int
	CMDTYPE    int
	RUN        int
	ALARMTYPE  int
	CMDNO      int
}

func (e Evem) createResponse() Evem {
	return Evem{
		MODULE:    "EVEM",
		OPERATION: e.OPERATION,
		PARAMETER: nil,
		SESSION:   e.SESSION,
		RESPONSE: &EvemResponse{
			SERIAL:     0, // 0: Release alarm   1: start the alarm   2：Pre alarm
			ERRORCAUSE: "SUCCESS",
			ERRORCODE:  0,
			ALARMUID:   e.PARAMETER.ALARMUID,
			CMDTYPE:    e.PARAMETER.CMDTYPE,
			RUN:        e.PARAMETER.RUN,
			ALARMTYPE:  e.PARAMETER.ALARMTYPE,
			CMDNO:      e.PARAMETER.CMDNO,
		},
	}
}

func (e Evem) HandleRequest(channel channel.IChannel, buffer []byte) {
	response := e.createResponse()
	responseJ, _ := json.Marshal(response)
	bytes := append(validMagicPackageHeader[:], responseJ...)
	err := channel.SendBytes(bytes)
	fmt.Printf("\nsent packet back as text: %s", bytes)
	if err != nil {
		fmt.Errorf("HandleRequest %e", err)
	}
}

func (e Evem) ParseDtoFromData(buffer []byte) interface{} {
	var result Evem
	jei := strings.LastIndex(string(buffer), "}")
	err := json.Unmarshal(buffer[12:jei+1], &result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}
