package dto

//"OPERATION":"MEDIATASKSTART"  ACK (callback)
type MediaStreamModelMediaTaskStartResponseParameter struct {
	CSRC       string
	IPANDPORT  string
	PT         int //payload type as in package header
	SSRC       int //128
	STREAMNAME string
}
