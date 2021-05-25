package dto

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
