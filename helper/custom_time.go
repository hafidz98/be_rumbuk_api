package helper

import (
	"time"
)

type CustomTime struct {
	time.Time
}

func (t *CustomTime) CusTime(b []uint8) error {
	formattedTime, err := time.Parse("15:04:05", string(b))
	t.Time = formattedTime
	return err
}