package dto

//"OPERATION":"MEDIATASKSTOP"  ACK (callback)
type MediaStreamModelMediaTaskStopResponseParameter struct {
	MediaStreamModelMediaTaskStartResponseParameter
	ERRORCAUSE string
	ERRORCODE  int
}
