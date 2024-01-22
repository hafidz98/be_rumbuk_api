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
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type RoomTimeSlot struct {
	RoomTimeSlotID int
	TimeSlotID     int
	RoomID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Rooms struct {
	Building Building
	Floor    Floor
	Room     Room
	TimeSlot TimeSlot
	Reserved *bool
}
