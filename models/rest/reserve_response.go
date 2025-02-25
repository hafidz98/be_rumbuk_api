package rest

import (
	"time"
)

type ReservationDetailResponse struct {
	ID           int       `json:"reservation_id"`
	BookDate     time.Time `json:"booking_date"`
	StudentID    string    `json:"student_id"`
	Activity     string    `json:"activity"`
	Status       string    `json:"status"`
	StudentName  string    `json:"student_name"`
	StartTime    string    `json:"start_time"`
	EndTime      string    `json:"end_time"`
	RoomName     string    `json:"room"`
	FloorName    string    `json:"floor"`
	BuildingName string    `json:"building"`
}

type ReserveResponse struct {
	ReserveID  int       `json:"id"`
	BookDate   time.Time `json:"booking_date"`
	StudentID  string    `json:"student_id"`
	Activity   string    `json:"activity"`
	Room       *RoomData `json:"room"`
	Status     string    `json:"status"`
	StatusText string    `json:"status_text,omitempty"`
}

type RoomData struct {
	RoomName     string `json:"room"`
	FloorName    string `json:"floor"`
	BuildingName string `json:"building"`
	Capacity     int    `json:"capacity"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}
