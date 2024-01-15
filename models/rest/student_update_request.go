package rest

type StudentUpdateRequest struct {
	StudentID   string `validate:"required" json:"student_id"`
	Name        string `validate:"required,max=255" json:"name"`
	Gender      string `json:"gender"`
	BatchYear   int    `json:"batch_year"`
	Major       string `json:"major"`
	Faculty     string `json:"faculty"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
