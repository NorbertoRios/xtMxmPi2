package dto

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
