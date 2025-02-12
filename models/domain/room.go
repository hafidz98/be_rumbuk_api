package domain

import (
	"time"
)

type Room struct {
	ID         int
	Name       string
	Capacity   int
	BuildingID int
	FloorID    int
	Status     string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	TimeSlot   []TimeSlot
}
