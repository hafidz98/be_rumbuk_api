package rest

type RoomTimeslotRequest struct {
	IDRoom      int    `json:"room_id"`
	RoomName    string `json:"room_name,omitempty"`
	TimeSlotIDs []int  `json:"time_slot_id"`
}
