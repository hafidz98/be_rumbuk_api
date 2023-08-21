package rest

type StaffResponse struct {
	StaffID string `json:"staff_id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Email   string `json:"email"`
}
