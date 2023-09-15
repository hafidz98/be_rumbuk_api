package rest

type BuildingCreateRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type BuildingUpdateRequest struct{
	ID       int    `json:"id"`
	Name     string `json:"name"`
}
