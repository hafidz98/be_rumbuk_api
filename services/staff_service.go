package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	service_model "github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type StaffService interface {
	Create(context context.Context, request service_model.StaffCreateRequest) service_model.StaffResponse
	Update(context context.Context, request service_model.StaffUpdateRequest) service_model.StaffResponse
	Delete(context context.Context, staffId string)
	FetchById(context context.Context, staffId string) service_model.StaffResponse
	FetchAllFilter(context context.Context, filter *domain.FilterParams) []service_model.StaffResponse
	Pagination(context context.Context, params *domain.FilterParams) (service_model.PaginationMeta, interface{})
}

type StaffServiceImpl struct {
	StaffRepository repositories.StaffRepo
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewStaffService(staffRepository repositories.StaffRepo, DB *sql.DB, validate *validator.Validate) StaffService {
	return &StaffServiceImpl{
		StaffRepository: staffRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func toStaffResponse(staff domain.Staff) service_model.StaffResponse {
	return service_model.StaffResponse{
		StaffID: staff.StaffID,
		Name:    staff.Name,
		Role:    staff.Role,
		Email:   staff.Email,
	}
}

func toStaffResponses(staffs []domain.Staff) []service_model.StaffResponse {
	var StaffResponses []service_model.StaffResponse
	for _, staff := range staffs {
		StaffResponses = append(StaffResponses, toStaffResponse(staff))
	}
	return StaffResponses
}

func (service *StaffServiceImpl) Create(context context.Context, request service_model.StaffCreateRequest) service_model.StaffResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staff := domain.Staff{
		StaffID:  request.StaffID,
		Name:     request.Name,
		Role:     request.Role,
		Email:    request.Email,
		Password: request.Password,
	}

	staff = service.StaffRepository.Create(context, tx, staff)
	return toStaffResponse(staff)
}

func (service *StaffServiceImpl) Update(context context.Context, request service_model.StaffUpdateRequest) service_model.StaffResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staff, err := service.StaffRepository.FetchById(context, tx, request.StaffID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	staff = domain.Staff{
		StaffID: request.StaffID,
		Name:    request.Name,
		Role:    request.Role,
		Status:  request.Status,
		Email:   request.Email,
	}

	staff = service.StaffRepository.Update(context, tx, staff)
	return toStaffResponse(staff)
}

func (service *StaffServiceImpl) Delete(context context.Context, staffId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staff, err := service.StaffRepository.FetchById(context, tx, staffId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	staff = domain.Staff{
		Status: "0",
	}

	service.StaffRepository.SoftDelete(context, tx, staff)
}

func (service *StaffServiceImpl) FetchById(context context.Context, staffId string) service_model.StaffResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staff, err := service.StaffRepository.FetchById(context, tx, staffId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return toStaffResponse(staff)
}

func (service *StaffServiceImpl) FetchAllFilter(context context.Context, filter *domain.FilterParams) []service_model.StaffResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staffs := service.StaffRepository.FetchAllFilter(context, tx, filter)
	return toStaffResponses(staffs)
}

func (service *StaffServiceImpl) Pagination(context context.Context, params *domain.FilterParams) (service_model.PaginationMeta, interface{}) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	staffCount := service.StaffRepository.CountAll(context, tx)

	helper.Warning.Printf("Record count: %v", staffCount)

	totalPage := int(math.Ceil(float64(staffCount) / float64(params.PerPage)))

	meta := &service_model.PaginationMeta{
		Page:    int(params.Page),
		PerPage: int(params.PerPage),
		Total:   totalPage,
	}

	links := make(map[string]string)
	links["self"] = fmt.Sprintf("?page=%d&per_page=%d", params.Page, params.PerPage)
	if params.Page > 1 {
		links["prev"] = fmt.Sprintf("?page=%d&per_page=%d", params.Page-1, params.PerPage)
	}
	if params.Page < uint64(totalPage) {
		links["next"] = fmt.Sprintf("?page=%d&per_page=%d", params.Page+1, params.PerPage)
	}
	links["last"] = fmt.Sprintf("?page=%d&per_page=%d", totalPage, params.PerPage)

	return *meta, links
}
