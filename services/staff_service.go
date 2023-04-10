package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	service_model "github.com/hafidz98/be_rumbuk_api/models/service"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type StaffService interface {
	FetchById(xontext context.Context, staffId string) service_model.StaffResponse
}

type StaffServiceImpl struct {
	StaffRepository repositories.StaffRepo
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewStaffService(staffRepository repositories.StaffRepo, DB *sql.DB, validate *validator.Validate) StaffService {
	return &StaffServiceImpl{
		StaffRepository: staffRepository,
		DB: DB,
		Validate: validate,
	}
}

func ToStaffResponse(staff domain.Staff) service_model.StaffResponse{
	return service_model.StaffResponse{
		StaffID: staff.StaffID,
		Name: staff.Name,
		Role: staff.Role,
		Email: staff.Email,
	}
}

func (service *StaffServiceImpl) FetchById(context context.Context, staffId string) service_model.StaffResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staff, err := service.StaffRepository.FetchById(context, tx, staffId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToStaffResponse(staff)
}
