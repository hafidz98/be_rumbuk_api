package rest

type TimeSlotResponse struct {
	ID        int    `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Duration  int    `json:"duration"`
	Reserved  *bool    `json:"reserved,omitempty"`
}
