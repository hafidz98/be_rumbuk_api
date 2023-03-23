package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/model/domain"
	"github.com/hafidz98/be_rumbuk_api/model/service"
	"github.com/hafidz98/be_rumbuk_api/repository"
)

// data model request dan data model response
// Logic atau bussiness intelligence

func ToStudentResponse(student domain.Students) service.StudentResponse {
	return service.StudentResponse{
		StudentID: student.StudentID,
		Name:      student.Name,
	}
}

func ToStudentResponses(students []domain.Students) []service.StudentResponse {
	var studentResponses []service.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, ToStudentResponse(student))
	}
	return studentResponses
}

type StudentService interface {
	Create(context context.Context, request service.StudentCreateRequest) service.StudentResponse
	Update(context context.Context, request service.StudentUpdateRequest) service.StudentResponse
	Delete(context context.Context, studentId string)
	FetchById(context context.Context, studentId string) service.StudentResponse
	FindAll(context context.Context) []service.StudentResponse
}

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepo
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewStudentService(studentRepository repository.StudentRepo, DB *sql.DB, validate *validator.Validate) StudentService {
	return &StudentServiceImpl{
		StudentRepository: studentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *StudentServiceImpl) Create(context context.Context, request service.StudentCreateRequest) service.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student := domain.Students{
		Name: request.Name,
	}

	student = service.StudentRepository.Create(context, tx, student)

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) Update(context context.Context, request service.StudentUpdateRequest) service.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, request.StudentID)
	helper.PanicIfError(err)

	student.Name = request.Name

	student = service.StudentRepository.Update(context, tx, student)

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) Delete(context context.Context, studentId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, studentId)
	helper.PanicIfError(err)

	service.StudentRepository.Delete(context, tx, student)
}

//Get and/or show data

func (service *StudentServiceImpl) FetchById(context context.Context, studentId string) service.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, studentId)
	helper.PanicIfError(err)

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) FindAll(context context.Context) []service.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := service.StudentRepository.FetchAll(context, tx)

	return ToStudentResponses(students)
}
