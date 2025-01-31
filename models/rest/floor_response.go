package rest

type FloorResponse struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	BuildingId   int            `json:"building_id"`
	BuildingName string         `json:"building_name,omitempty"`
	Rooms        []RoomResponse `json:"rooms,omitempty"`
}
