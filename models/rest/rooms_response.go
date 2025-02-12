package rest

import (
	"time"
)

type RoomResponse struct {
	ID        int                `json:"id"`
	Name      string             `json:"name"`
	Capacity  int                `json:"capacity"`
	Building  int                `json:"building,omitempty"`
	Floor     int                `json:"floor,omitempty"`
	Status    string             `json:"status,omitempty"`
	CreatedAt *time.Time         `json:"created_at,omitempty"`
	UpdatedAt *time.Time         `json:"updated_at,omitempty"`
	TimeSlot  []TimeSlotResponse `json:"time_slot,omitempty"`
	Reserved  bool               `json:"reserved,omitempty"`
}
