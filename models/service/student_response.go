package service

type StudentResponse struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	BatchYear   int    `json:"batch_year"`
	Major       string `json:"major"`
	Faculty     string `json:"faculty"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
