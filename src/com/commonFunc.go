package com

import (
	"strconv"
	"time"
)

//14 digits input
//"20210526080429",
func N9MTimeToTime(n9mTime string) time.Time {
	y, _ := strconv.Atoi(n9mTime[:4])
	m, _ := strconv.Atoi(n9mTime[4:6])
	d, _ := strconv.Atoi(n9mTime[6:8])
	h, _ := strconv.Atoi(n9mTime[8:10])
	min, _ := strconv.Atoi(n9mTime[10:12])
	s, _ := strconv.Atoi(n9mTime[12:14])
	return time.Date(y, time.Month(m), d, h, min, s, 0, nil)
}
