package rest

type FloorResponse struct {
	ID     int     `json:"id"`
	Number string  `json:"number"`
	Rooms  []Rooms `json:"rooms"`
}
