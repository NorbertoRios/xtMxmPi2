package dto

type GetCalendarResponse struct {
	CALENDER   []string
	CHANNEL    int
	CHCALENDER []map[string]interface{}
	COUNT      int
	ERRORCAUSE string
	ERRORCODE  int
}
