package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
)

type Evem struct {
	GeneralPackageHeader `json:"-"`
	MODULE               string
	OPERATION            string
	PARAMETER            *EvemParameter `json:",omitempty"`
	SESSION              string
	RESPONSE             *EvemResponse `json:",omitempty"`
}

type AlarmType int
type AlarmImportance int
type Lang int

const (
	VideoLossAT AlarmType = iota
	CameraCoveredAT
	MotionDetectionAT
	StorageAbnormalAT
	UserDefinedAT
	SentriesInspectionAT
	ViolationDetectionAT
	EmergencyAT
	SpeedAT
	LowVoltageAT
	InternalUseAT        AlarmType = 16
	FenceInOutAT         AlarmType = 17
	AccAT                AlarmType = 18
	PeripheralsDroppedAT AlarmType = 19
	StopAnnouncementAT   AlarmType = 20
	GPSAntennaAT         AlarmType = 21
	DayNightSwitchAT     AlarmType = 22
	ProhibitDrivingAT    AlarmType = 23
	//24-31 reserved
	SerialAlarmAT                    AlarmType = 32
	FatigueAT                        AlarmType = 33
	TimeOutParkingAT                 AlarmType = 34
	GestureAlarmAT                   AlarmType = 35
	GreenDrivingAlarmEventAT         AlarmType = 36
	IllegalIgnitionAT                AlarmType = 37
	IllegalShutdownAT                AlarmType = 38
	CustomExternalInputAlarmAT       AlarmType = 39
	OilAlarmAT                       AlarmType = 42
	BusLaneOccupationAlarmAT         AlarmType = 43
	ForgottenAlarmAT                 AlarmType = 44
	SpecialCustomerFaultAlarmAT      AlarmType = 45
	TemperatureAbnormalAlarmAT       AlarmType = 46
	TemperatureChangeAbnormalAlarmAT AlarmType = 47
	SmokeAlarmAT                     AlarmType = 48
	GBoxAlarmAT                      AlarmType = 49
	LicensePlateRecognitionAlarmAT   AlarmType = 50
	SpeedAlarmAT                     AlarmType = 51
	WirelessSignalAbnormalAlarmAT    AlarmType = 52
	ArmingAlarmAT                    AlarmType = 53
	PhoneCallAlarm                   AlarmType = 54
	GPSFaultAlarm                    AlarmType = 55
	DSMAlarmAT                       AlarmType = 56
	FireBoxAlarm                     AlarmType = 57
	DriverFacialRecognitionAlarm     AlarmType = 96
)

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

type EvemParameter struct {
	AlarmType       AlarmType       `json:"ALARMTYPE,omitempty"`
	AlarmImportance AlarmImportance `json:"ALARMAS,omitempty"`
	AlarmCount      uint16          `json:"ALARMCOUNT,omitempty"`
	AlarmUID        int             `json:"ALARMUID,omitempty"`
	CMDNo           int             `json:"CMDNO,omitempty"`
	CMDType         int             `json:"CMDTYPE,omitempty"`
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
	RUN         int //Consistent with the reporting field
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
			ALARMUID:   e.PARAMETER.AlarmUID,
			CMDTYPE:    e.PARAMETER.CMDType,
			RUN:        e.PARAMETER.RUN,
			ALARMTYPE:  e.PARAMETER.AlarmType,
			CMDNO:      e.PARAMETER.CMDNo,
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
	e.FillGeneralPackageHeaderFromPackage(buffer)
	err := json.Unmarshal(e.PayloadBody, &result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}
