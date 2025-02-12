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

type TimeslotService interface {
	Create(context context.Context, request rest.TimeSlotCreateRequest) rest.TimeSlotResponse
	Update(context context.Context, request rest.TimeSlotUpdateRequest) rest.TimeSlotResponse
	//Delete(context context.Context, timeslotId int)
	//GetById(context context.Context, timeslotId int) []rest.RoomResponse
	GetById(context context.Context, timeslotId int) rest.TimeSlotResponse
	GetAll(context context.Context) []rest.TimeSlotResponse
}

type TimeslotServiceImpl struct {
	TimeSlotRepo repositories.TimeSlotRepo
	DB           *sql.DB
	Validate     *validator.Validate
}

func NewTimeslotService(timeslotRepository repositories.TimeSlotRepo, DB *sql.DB, validate *validator.Validate) TimeslotService {
	return &TimeslotServiceImpl{
		TimeSlotRepo: timeslotRepository,
		DB:           DB,
		Validate:     validate,
	}
}

//service composite

func (service *TimeslotServiceImpl) Create(context context.Context, request rest.TimeSlotCreateRequest) rest.TimeSlotResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	timeslot := domain.TimeSlot{
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Duration:  request.Duration,
	}

	timeslot = service.TimeSlotRepo.Create(context, tx, timeslot)

	return ToTimeslotResponse(timeslot)
}

func (service *TimeslotServiceImpl) Update(context context.Context, request rest.TimeSlotUpdateRequest) rest.TimeSlotResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	timeslot, err := service.TimeSlotRepo.SelectById(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	timeslot = domain.TimeSlot{
		StartTime: request.StartTime,
		EndTime:   request.EndTime,
		Duration:  request.Duration,
	}

	timeslot = service.TimeSlotRepo.Update(context, tx, timeslot)

	return ToTimeslotResponse(timeslot)
}

//func (service *TimeslotServiceImpl) Delete(context context.Context, timeslotId int)

//func (service *TimeslotServiceImpl) GetById(context context.Context, timeslotId int) []rest.RoomResponse

func (service *TimeslotServiceImpl) GetById(context context.Context, timeslotId int) rest.TimeSlotResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	timeslot, err := service.TimeSlotRepo.SelectById(context, tx, timeslotId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToTimeslotResponse(timeslot)
}

func (service *TimeslotServiceImpl) GetAll(context context.Context) []rest.TimeSlotResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	timeslot := service.TimeSlotRepo.SelectAll(context, tx)

	return ToTimeslotResponses(timeslot)
}

func ToTimeslotResponse(timeslot domain.TimeSlot) rest.TimeSlotResponse {
	return rest.TimeSlotResponse{
		ID:        timeslot.ID,
		StartTime: timeslot.StartTime,
		EndTime:   timeslot.EndTime,
		Duration:  timeslot.Duration,
	}
}

func ToTimeslotResponses(timeslots []domain.TimeSlot) []rest.TimeSlotResponse {
	var timeslotResponses []rest.TimeSlotResponse
	for _, timeslot := range timeslots {
		timeslotResponses = append(timeslotResponses, ToTimeslotResponse(timeslot))
	}
	return timeslotResponses
}
