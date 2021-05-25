package dto

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
