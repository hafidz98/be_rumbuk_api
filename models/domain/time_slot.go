package domain

import (
	"time"
)

type TimeSlot struct {
	ID        int
	StartTime string
	EndTime   string
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
