package rest

type RoomCreateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Building int    `json:"building"`
	Floor    int    `json:"floor"`
}

type RoomUpdateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Capacity int    `json:"capacity,omitempty"`
}
