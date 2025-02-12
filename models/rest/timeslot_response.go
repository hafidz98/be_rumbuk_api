package rest

type TimeSlotResponse struct {
	RoomTSID  int    `json:"rts_id,omitempty"`
	ID        int    `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  int    `json:"duration,omitempty"`
	Reserved  *bool  `json:"reserved,omitempty"`
}
