package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	service_model "github.com/hafidz98/be_rumbuk_api/models/service"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type AuthService interface {
	Login(context context.Context, request service_model.AuthLoginRequest) (tokenString string)
}

type AuthServiceImpl struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func NewAuthService(DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		DB:       DB,
		Validate: validate,
	}
}

func (service *AuthServiceImpl) Login(context context.Context, request service_model.AuthLoginRequest) (tokenString string) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userStaff, err := repositories.NewStaffRepo().FetchById(context, tx, request.UserID)
	match := helper.ComparePassword(userStaff.Password, request.Password)

	if err == nil && match {
		userData := service_model.GlobalJWTResponse{
			UserID: userStaff.StaffID,
			Role:   "Staff",
		}

		claims := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		token, err := helper.GenerateJWT(&userData, claims)
		helper.PanicIfError(err)

		return token
	} else if err == nil && !match {
		userStudent, err := repositories.NewStudentRepo().FetchBySId(context, tx, request.UserID)
		match := helper.ComparePassword(userStudent.Password, request.Password)
		if err != nil || !match {
			panic(exception.NewAuthorization(exception.InvalidCredentials))
		}

		userData := service_model.GlobalJWTResponse{
			UserID: userStudent.StudentID,
			Name:   userStudent.Name,
			Email:  userStudent.Email,
			Role:   "Student",
		}

		claims := jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}

		token, err := helper.GenerateJWT(&userData, claims)
		helper.PanicIfError(err)

		return token
	}
	
	panic(exception.NewAuthorization(exception.InvalidCredentials))
}
