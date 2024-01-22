package rest

type FloorResponse struct {
	ID     int     `json:"id"`
	Name string  `json:"name"`
	Rooms  []RoomResponse `json:"rooms,omitempty"`
}
