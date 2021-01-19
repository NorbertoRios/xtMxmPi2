package dto

import (
	"comm/channel"
	"encoding/json"
	"fmt"
	"reflect"
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
	PhoneCallAlarmAT                 AlarmType = 54
	GPSFaultAlarm                    AlarmType = 55
	DSMAlarmAT                       AlarmType = 56 //phone not allowed alarm
	FireBoxAlarm                     AlarmType = 57
	DriverFacialRecognitionAlarm     AlarmType = 96
)

var AlarmTypeMap = map[int]reflect.Type{
	0:  reflect.TypeOf(ChannelNumberAlarmParameter{}),
	1:  reflect.TypeOf(ChannelNumberAlarmParameter{}),
	2:  reflect.TypeOf(ChannelNumberAlarmParameter{}),
	3:  reflect.TypeOf(MemoryAbnormalAlarmParameter{}),
	4:  reflect.TypeOf(UserDefinedAlarmParameter{}),
	5:  reflect.TypeOf(SentryInspectionAlarmParameter{}),
	6:  reflect.TypeOf(ChannelNumberAlarmParameter{}),
	7:  reflect.TypeOf(EmergencyAlarmParameter{}),
	8:  reflect.TypeOf(SpeedAlarmParameter{}),
	9:  reflect.TypeOf(LowVoltageAlarmParameter{}),
	17: reflect.TypeOf(GeoFenceAlarmParameter{}),
	18: reflect.TypeOf(AccAlarmParameter{}),
	19: reflect.TypeOf(PeripheralDroppedAlarmParameter{}),
	20: reflect.TypeOf(StopAnnouncementAlarmParameter{}),
	21: reflect.TypeOf(GPSAntennaAlarmParameter{}),
	22: reflect.TypeOf(DayNightSwitchAlarm{}),
	32: reflect.TypeOf(SerialAlarmParameter{}),
	33: reflect.TypeOf(FatigueDrivingAlarmParameter{}),
	34: reflect.TypeOf(TimeoutParkingAlarmParameter{}),
	35: reflect.TypeOf(GestureAlarmParameter{}),
	36: reflect.TypeOf(GreenDrivingAlarmParameter{}),
	37: reflect.TypeOf(IllegalIgnitionAlarm{}),
	38: reflect.TypeOf(IllegalShutdownAlarm{}),
	39: reflect.TypeOf(CustomExternalInputAlarm{}),
	42: reflect.TypeOf(OilVolumeAlarmParameter{}),
	43: reflect.TypeOf(BusLaneOccupationAlarmParameter{}),
	44: reflect.TypeOf(UserDefinedAlarmParameter{}),
	45: reflect.TypeOf(SpecialCustomerMalfunctionAlarmParameter{}),
	46: reflect.TypeOf(TemperatureAbnormallyAlarmParameter{}),
	47: reflect.TypeOf(AbnormalTemperatureChangeAlarmParameter{}),
	48: reflect.TypeOf(SmokeAlarmParameter{}),
	49: reflect.TypeOf(GBoxAlarmParameter{}),
	50: reflect.TypeOf(LicensePlateRecognitionAlarmParameter{}),
	51: reflect.TypeOf(SpeedAlarmParameter{}),
	52: reflect.TypeOf(WirelessSignalAbnormalityAlarmParameter{}),
	53: reflect.TypeOf(ArmingAlarmParameter{}),
	54: reflect.TypeOf(PhoneCallAlarm{}),
	55: reflect.TypeOf(GPSMalfunctionAlarmParameter{}),
	56: reflect.TypeOf(DSMAlarmParameter{}),
	57: reflect.TypeOf(FireBoxAlarmParameter{}),
}

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
