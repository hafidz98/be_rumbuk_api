package domain

import (
	"time"
	//"github.com/hafidz98/be_rumbuk_api/utils"
)

// Reservation Status
//
// 0: Delete (Canceled);
// 1: Coming Soon; Upcoming;
// 2: Ongoing;
// 3: Done
type Reservation struct {
	ID             int
	BookDate       time.Time
	StudentID      string
	Activity       string
	RoomTimeSlotID int
	Status         string
}
