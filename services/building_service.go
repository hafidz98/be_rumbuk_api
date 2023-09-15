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

func ToBuildingResponse(building domain.Building) rest.BuildingResponse {
	return rest.BuildingResponse{
		ID:        building.ID,
		Name:      building.Name,
		CreatedAt: building.CreatedAt,
		UpdatedAt: building.UpdatedAt,
	}
}

func ToBuildingResponses(building []domain.Building) []rest.BuildingResponse {
	var buildingResponses []rest.BuildingResponse
	for _, building := range building {
		buildingResponses = append(buildingResponses, ToBuildingResponse(building))
	}
	return buildingResponses
}

type BuildingService interface {
	Create(context context.Context, request rest.BuildingCreateRequest) rest.BuildingResponse
	Update(context context.Context, request rest.BuildingUpdateRequest) rest.BuildingResponse
	Delete(context context.Context, buildingId int)
	FetchById(context context.Context, buildingId int) rest.BuildingResponse
	FetchAll(context context.Context) []rest.BuildingResponse
}

type BuildingServiceImpl struct {
	BuildingRepository repositories.BuildingRepo
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewBuildingService(buildingRepository repositories.BuildingRepo, DB *sql.DB, validate *validator.Validate) BuildingService {
	return &BuildingServiceImpl{
		BuildingRepository: buildingRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *BuildingServiceImpl) Create(context context.Context, request rest.BuildingCreateRequest) rest.BuildingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building := domain.Building{
		Name: request.Name,
	}

	building = service.BuildingRepository.Create(context, tx, building)

	return ToBuildingResponse(building)
}

func (service *BuildingServiceImpl) Update(context context.Context, request rest.BuildingUpdateRequest) rest.BuildingResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FetchByID(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	building = domain.Building{
		Name: request.Name,
	}

	building = service.BuildingRepository.Update(context, tx, building)

	return ToBuildingResponse(building)
}

func (service *BuildingServiceImpl) Delete(context context.Context, buildingId int) {}

func (service *BuildingServiceImpl) FetchById(context context.Context, buildingId int) rest.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building, err := service.BuildingRepository.FetchByID(context, tx, buildingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToBuildingResponse(building)
}

func (service *BuildingServiceImpl) FetchAll(context context.Context) []rest.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	building := service.BuildingRepository.FetchAll(context, tx)

	return ToBuildingResponses(building)
}
