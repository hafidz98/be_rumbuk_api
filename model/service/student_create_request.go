package service

// representasi dari request

type StudentCreateRequest struct {
	Name string `validate:"required,max=255" json:"name"`
}
