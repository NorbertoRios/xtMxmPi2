package dto

type QueryFileListRequestParameter struct {
	SERIAL     uint //unique request identifier
	STARTTIME  string
	CHANNEL    int
	ENDTIME    string
	STREAMTYPE int
	FILETYPE   int
	RFSTORAGE  int // 0 hard drive, 1 sd card
}
