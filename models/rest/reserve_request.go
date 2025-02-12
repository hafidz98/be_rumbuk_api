package rest

type ReserveCreateRequest struct {
	BookDate       string `json:"booking_date"`
	StudentID      string `json:"student_id"`
	Activity       string `json:"activity"`
	RoomTimeSlotID int    `json:"room_timeslot_id"`
	RoomID         int    `json:"room_id"`
	TimeslotID     int    `json:"timeslot_id"`
}

type ReserveUpdateRequest struct {
	ReservationID int `json:"reservation_id"`
	Status        int `json:"status,omitempty"`
}
