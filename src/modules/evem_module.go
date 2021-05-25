package modules

import (
	"encoding/json"
	"fmt"
	"streamax-go/dto"
	"streamax-go/interfaces"
)

type Evem struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            *interface{} `json:",omitempty"`
	SESSION              string
	RESPONSE             *dto.EvemResponse `json:",omitempty"`
	alarmType            int
	marshalParam         []byte
}

func (e Evem) createResponse(eP *dto.ActualDSMAlarmParameter) Evem {
	return Evem{
		MODULE:    "EVEM",
		OPERATION: e.OPERATION,
		PARAMETER: nil,
		SESSION:   e.SESSION,
		RESPONSE: &dto.EvemResponse{
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
		var p dto.ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 1:
		var p dto.ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 2:
		var p dto.ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 3:
		var p dto.MemoryAbnormalAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 4:
		var p dto.UserDefinedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 5:
		var p dto.SentryInspectionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 6:
		var p dto.ChannelNumberAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 7:
		var p dto.EmergencyAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 8:
		var p dto.SpeedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 9:
		var p dto.LowVoltageAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 17:
		var p dto.GeoFenceAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 18:
		var p dto.AccAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 19:
		var p dto.PeripheralDroppedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 20:
		var p dto.StopAnnouncementAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 21:
		var p dto.GPSAntennaAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 22:
		var p dto.DayNightSwitchAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 32:
		var p dto.SerialAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 33:
		var p dto.FatigueDrivingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 34:
		var p dto.TimeoutParkingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 35:
		var p dto.GestureAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 36:
		var p dto.GreenDrivingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 37:
		var p dto.IllegalIgnitionAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 38:
		var p dto.IllegalShutdownAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 39:
		var p dto.CustomExternalInputAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 42:
		var p dto.OilVolumeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 43:
		var p dto.BusLaneOccupationAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 44:
		var p dto.UserDefinedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 45:
		var p dto.SpecialCustomerMalfunctionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 46:
		var p dto.TemperatureAbnormallyAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 47:
		var p dto.AbnormalTemperatureChangeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 48:
		var p dto.SmokeAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 49:
		var p dto.GBoxAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 50:
		var p dto.LicensePlateRecognitionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 51:
		var p dto.SpeedAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 52:
		var p dto.WirelessSignalAbnormalityAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 53:
		var p dto.ArmingAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 54:
		var p dto.PhoneCallAlarm
		json.Unmarshal(e.marshalParam, &p)
	case 55:
		var p dto.GPSMalfunctionAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
	case 56:
		//var p DSMAlarmParameter
		var p dto.ActualDSMAlarmParameter
		json.Unmarshal(e.marshalParam, &p)
		response := e.createResponse(&p)
		responseJ, _ := json.Marshal(response)
		return responseJ
	case 57:
		var p dto.FireBoxAlarmParameter
		json.Unmarshal(e.marshalParam, &p)

	}
	return []byte{}
}
