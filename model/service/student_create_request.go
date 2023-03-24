package service

// representasi dari request

type StudentCreateRequest struct {
	StudentID string `validate:"required" json:"student_id"`
	Name      string `validate:"required,max=255" json:"name"`
}
