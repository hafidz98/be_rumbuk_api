package helper

import "time"

func FormatDate(date string) time.Time {
	timeFormat := "2006-01-02"
	t, _ := time.Parse(timeFormat, date)
	return t.Local()
}