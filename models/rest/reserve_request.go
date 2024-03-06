package rest

import (
	"time"
	//"github.com/hafidz98/be_rumbuk_api/utils"
)//

type ReserveCreateRequest struct {
	BookDate       time.Time `json:"booking_date"`
	StudentID      string    `json:"student_id"`
	Activity       string    `json:"activity"`
	RoomTimeSlotID int       `json:"room_timeslot_id"`
}
