package domain

import (
	"time"
)

type Reservation struct {
	ReservationID int
	StudentID     string
	RsvDate       time.Time
	Activity      string
	Status        string
}

type RoomReserved struct {
	RoomReservedID int
	ReservationID  int
	RoomTimeSlotID int
}
