package rest

import (
	"time"
)

type ReserveCreateRequest struct {
	BookDate       time.Time `json:"booking_date"`
	StudentID      string    `json:"student_id"`
	Activity       string    `json:"activity"`
	RoomTimeSlotID int       `json:"room_timeslot_id"`
}

type ReserveUpdateRequest struct {
	ReservationID int `json:"reservation_id"`
	Status        int `json:"status,omitempty"`
}
