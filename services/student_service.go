package services

import (
	"context"
	"database/sql"
	//"time"

	//"errors"

	"github.com/go-playground/validator/v10"
	//"github.com/golang-jwt/jwt/v5"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	service_model "github.com/hafidz98/be_rumbuk_api/models/service"
	repository "github.com/hafidz98/be_rumbuk_api/repositories"
)

// data model request dan data model response
// Logic atau bussiness intelligence

func ToStudentResponse(student domain.Students) service_model.StudentResponse {
	return service_model.StudentResponse{
		StudentID:   student.StudentID,
		Name:        student.Name,
		Gender:      student.Gender,
		BatchYear:   student.BatchYear,
		Major:       student.Major,
		Faculty:     student.Faculty,
		PhoneNumber: student.PhoneNumber,
		Email:       student.Email,
	}
}

func ToStudentResponses(students []domain.Students) []service_model.StudentResponse {
	var studentResponses []service_model.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, ToStudentResponse(student))
	}
	return studentResponses
}

type StudentService interface {
	Create(context context.Context, request service_model.StudentCreateRequest) service_model.StudentResponse
	Update(context context.Context, request service_model.StudentUpdateRequest) service_model.StudentResponse
	Delete(context context.Context, studentId string)
	FetchById(context context.Context, studentId string) service_model.StudentResponse
	FindAll(context context.Context) []service_model.StudentResponse
	//Login(context context.Context, request service_model.StudentLoginRequest) (tokenString string)
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

// func (service *StudentServiceImpl) Login(context context.Context, request service_model.StudentLoginRequest) (tokenString string) {
// 	err := service.Validate.Struct(request)
// 	helper.PanicIfError(err)

// 	tx, err := service.DB.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	student, err := service.StudentRepository.FetchBySId(context, tx, request.StudentID)
// 	if err != nil {
// 		panic(exception.NewAuthorization(exception.InvalidCredentials))
// 	}

// 	match := helper.ComparePassword(student.Password, request.Password)
// 	if !match {
// 		panic(exception.NewAuthorization(exception.InvalidCredentials))
// 	}

// 	s := service_model.GlobalJWTResponse{
// 		UserID: student.StudentID,
// 		Name:   student.Name,
// 		Email:  student.Email,
// 		Role:   "Student",
// 	}

// 	claims := jwt.RegisteredClaims{
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
// 	}

// 	token, err := helper.GenerateJWT(&s, claims)
// 	helper.PanicIfError(err)

// 	helper.Info.Println("Token: " + token)
// 	helper.Info.Println("Login request granted")
// 	return token
// }

func (service *StudentServiceImpl) Create(context context.Context, request service_model.StudentCreateRequest) service_model.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	password, err := helper.GenerateHashedPassword(request.Password)
	helper.PanicIfError(err)

	student := domain.Students{
		StudentID:   request.StudentID,
		Name:        request.Name,
		Gender:      request.Gender,
		BatchYear:   request.BatchYear,
		Major:       request.Major,
		Faculty:     request.Faculty,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
		Password:    password,
	}

	student = service.StudentRepository.Create(context, tx, student)

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) Update(context context.Context, request service_model.StudentUpdateRequest) service_model.StudentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, request.StudentID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//student.Name = request.Name

	student = domain.Students{
		StudentID:   request.StudentID,
		Name:        request.Name,
		Gender:      request.Gender,
		BatchYear:   request.BatchYear,
		Major:       request.Major,
		Faculty:     request.Faculty,
		PhoneNumber: request.PhoneNumber,
		Email:       request.Email,
	}

	student = service.StudentRepository.Update(context, tx, student)

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) Delete(context context.Context, studentId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.StudentRepository.Delete(context, tx, student)
}

//Get and/or show data

func (service *StudentServiceImpl) FetchById(context context.Context, studentId string) service_model.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	student, err := service.StudentRepository.FetchBySId(context, tx, studentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToStudentResponse(student)
}

func (service *StudentServiceImpl) FindAll(context context.Context) []service_model.StudentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	students := service.StudentRepository.FetchAll(context, tx)

	return ToStudentResponses(students)
}
