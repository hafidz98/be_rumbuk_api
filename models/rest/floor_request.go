package rest

type FloorCreateRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	BuildingID int    `json:"building_id"`
}
