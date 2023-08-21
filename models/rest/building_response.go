package rest

type BuildingResponse struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Floors []FloorResponse `json:"floors"`
}
