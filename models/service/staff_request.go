package service

type StaffCreateRequest struct {
	StaffID  string `validate:"required" json:"staff_id"`
	Name     string `validate:"required,max=255" json:"name"`
	Role     string `validate:"required,max=255" json:"role"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type StaffUpdateRequest struct {
	StaffID  string `validate:"required" json:"staff_id"`
	Name     string `validate:"required,max=255" json:"name"`
	Role     string `validate:"required,max=55" json:"role"`
	Status   string `validate:"required,max=2" json:"status"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
