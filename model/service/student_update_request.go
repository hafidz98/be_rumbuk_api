package service

type StudentUpdateRequest struct{
	StudentID string `validate:"required"`
	Name string `validate:"required,max=255"`
}