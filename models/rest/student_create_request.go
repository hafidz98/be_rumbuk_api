package rest

// representasi dari request

type StudentCreateRequest struct {
	StudentID   string `validate:"required" json:"student_id"`
	Name        string `validate:"required,max=255" json:"name"`
	Gender      string `validate:"required" json:"gender"`
	BatchYear   int    `validate:"required" json:"batch_year"`
	Major       string `validate:"required" json:"major"`
	Faculty     string `validate:"required" json:"faculty"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	Email       string `validate:"required" json:"email"`
	Password    string `validate:"required" json:"password"`
}
