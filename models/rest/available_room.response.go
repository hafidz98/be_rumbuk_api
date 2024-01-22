package rest

import "time"

type AvailabeRoomResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Floors    []FloorResponse `json:"floors,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
}