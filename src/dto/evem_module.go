package dto

import (
	"encoding/json"
	"fmt"
	"streamax-go/interfaces"
)

type Evem struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            *interface{} `json:",omitempty"`
	SESSION              string
	RESPONSE             *EvemResponse `json:",omitempty"`
	alarmType            int
	marshalParam         []byte
}

type AlarmType int
type AlarmImportance int
type Lang int

const (
	ImportantEventAI AlarmImportance = iota
	GeneralAlarmAI
	EmergencyAlarmAI
)

const (
	SimplifiedChinese Lang = iota
	English
	Korean
	Italian
	German
	Thai
	Turkey
	Portugal
	Spain
	Romania
	Greece
	French
	Russian
	Dutch
	Hebrew
	ChineseTraditional
)

type ActualDSMAlarmParameter struct {
	GeneralParamPayload
	AlarmImportance AlarmImportance `json:"ALARMAS,omitempty"`
	AlarmCount      uint16          `json:"ALARMCOUNT,omitempty"`
	CurrentTime     int             `json:"CURRENTTIME,omitempty"`
	EventUUID       string          `json:"EVTUUID,omitempty"`
	Language        int             `json:"L,omitempty"`
	LEV             int             `json:"LEV,omitempty"`
	P               *struct {       //gps object
		C int
		J string
		S int
		T string
		V int
		W string
	}
	TO          int //Indicates reported platform
	REAL        int // 0: means real-time upload; 1: Indicates replenishing uploading
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
	ALARMTYPE  AlarmType
	CMDNO      int
}

type GeneralParamPayload struct {
	ALARMUID int //Important distinctive counter, +1 for each new alarm, unique by alarm
	//0- Disable alarm, 1--Start alarm, 2-- Pre Alarm.   3-Continuous alarm
	//camera itself changing states, long alarms use CMDTYPE 1 and then same data with 0 and CMDNO +1
	//or timeless alarms use CMDTYPE 1 with no repeats
	CMDTYPE   int
	RUN       int
	ALARMTYPE AlarmType
	CMDNO     int //increasing in each command (start alarm or disable alarm via CMDTYPE)
}

func (e Evem) createResponse(eP *ActualDSMAlarmParameter) Evem {
	return Evem{
		MODULE:    "EVEM",
		OPERATION: e.OPERATION,
		PARAMETER: nil,
		SESSION:   e.SESSION,
		RESPONSE: &EvemResponse{
			SERIAL:     0, // 0: Release alarm   1: start the alarm   2：Pre alarm
			ERRORCAUSE: "SUCCESS",
			ERRORCODE:  0,
			ALARMUID:   eP.ALARMUID,
			CMDTYPE:    eP.CMDTYPE,
			RUN:        eP.RUN,
			ALARMTYPE:  eP.ALARMTYPE,
			CMDNO:      eP.CMDNO,
		},
	}
}

func (e Evem) HandleRequest(channel interfaces.IChannel, buffer []byte) {
	jBytes := e.callAlarmHandler()
	bytes := append(e.toHeaderBytes(uint(len(jBytes))), jBytes...)
	err := channel.SendBytes(bytes)
	fmt.Printf("\nsent packet back as text: %s", bytes)
	if err != nil {
		fmt.Errorf("HandleRequest %e", err)
	}
}

func (e Evem) ParseDtoFromData(buffer []byte) interface{} {
	var result Evem
	e.FillGeneralPackageHeaderFromPackage(buffer)
	var m map[string]interface{}
	json.Unmarshal(e.PayloadBody, &m)
	pm := m["PARAMETER"].(map[string]interface{})
	marshalParam, _ := json.Marshal(m["PARAMETER"])
	result.alarmType = int(pm["ALARMTYPE"].(float64))
	result.marshalParam = marshalParam
	return result
}

func (e *Evem) callAlarmHandler() []byte {
	switch e.alarmType {

	case 0:
		var p ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 1:
		var p ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 2:
		var p ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 3:
		var p MemoryAbnormalAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 4:
		var p UserDefinedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 5:
		var p SentryInspectionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 6:
		var p ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 7:
		var p EmergencyAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 8:
		var p SpeedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 9:
		var p LowVoltageAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 17:
		var p GeoFenceAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 18:
		var p AccAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 19:
		var p PeripheralDroppedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 20:
		var p StopAnnouncementAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 21:
		var p GPSAntennaAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 22:
		var p DayNightSwitchAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 32:
		var p SerialAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 33:
		var p FatigueDrivingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 34:
		var p TimeoutParkingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 35:
		var p GestureAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 36:
		var p GreenDrivingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 37:
		var p IllegalIgnitionAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 38:
		var p IllegalShutdownAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 39:
		var p CustomExternalInputAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 42:
		var p OilVolumeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 43:
		var p BusLaneOccupationAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 44:
		var p UserDefinedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 45:
		var p SpecialCustomerMalfunctionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 46:
		var p TemperatureAbnormallyAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 47:
		var p AbnormalTemperatureChangeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 48:
		var p SmokeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 49:
		var p GBoxAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 50:
		var p LicensePlateRecognitionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 51:
		var p SpeedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 52:
		var p WirelessSignalAbnormalityAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 53:
		var p ArmingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 54:
		var p PhoneCallAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 55:
		var p GPSMalfunctionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 56:
		//var p DSMAlarmParameter
		var p ActualDSMAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
		response := e.createResponse(&p)
		responseJ, _ := json.Marshal(response)
		return responseJ
	case 57:
		var p FireBoxAlarmParameter
		json.Unmarshal(e.marshalParam, &p)

	}
	return []byte{}
}
