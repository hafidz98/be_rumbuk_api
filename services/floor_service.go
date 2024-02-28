package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type FloorService interface {
	Create(context context.Context, request rest.FloorCreateRequest) rest.FloorResponse
	GetById(context context.Context, floorId int) rest.FloorResponse
	GetAll(context context.Context) []rest.FloorResponse
}

type FloorServiceImpl struct {
	FloorRepo repositories.FloorRepo
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewFloorService(floorRepository repositories.FloorRepo, DB *sql.DB, validate *validator.Validate) FloorService {
	return &FloorServiceImpl{
		FloorRepo: floorRepository,
		DB:        DB,
		Validate:  validate,
	}
}

func (service *FloorServiceImpl) Create(context context.Context, request rest.FloorCreateRequest) rest.FloorResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor := domain.Floor{
		ID:         request.ID,
		Name:       request.Name,
		BuildingID: request.BuildingID,
	}

	floor = service.FloorRepo.Create(context, tx, floor)

	return ToFloorResponse(floor)
}

func (service *FloorServiceImpl) GetById(context context.Context, floorId int) rest.FloorResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor, err := service.FloorRepo.SelectById(context, tx, floorId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToFloorResponse(floor)
}

func (service *FloorServiceImpl) GetAll(context context.Context) []rest.FloorResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	floor := service.FloorRepo.SelectAll(context, tx)

	return ToFloorResponses(floor)
}

func ToFloorResponse(floor domain.Floor) rest.FloorResponse {
	return rest.FloorResponse{
		ID:   floor.ID,
		Name: floor.Name,
	}
}

func ToFloorResponses(floors []domain.Floor) []rest.FloorResponse {
	var floorResponses []rest.FloorResponse
	for _, floor := range floors {
		floorResponses = append(floorResponses, ToFloorResponse(floor))
	}
	return floorResponses
}
