package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type RoomTimeslotService interface {
	AddRoomTimeslot(context context.Context, roomTimeslotReq rest.RoomTimeslotRequest) (rest.RoomTimeslotRequest, error)
}

type RoomTimeslotServiceImpl struct {
	RoomTimeslotRepository repositories.RoomTimeslotRepo
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewRoomTimeslotService(roomTimeslotRepository repositories.RoomTimeslotRepo, DB *sql.DB, validate *validator.Validate) RoomTimeslotService {
	return &RoomTimeslotServiceImpl{
		RoomTimeslotRepository: roomTimeslotRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *RoomTimeslotServiceImpl) AddRoomTimeslot(context context.Context, roomTimeslotReq rest.RoomTimeslotRequest) (rest.RoomTimeslotRequest, error) {
	err := service.Validate.Struct(roomTimeslotReq)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roomTimeslot := domain.RoomTimeslot{
		IDRoom:      roomTimeslotReq.IDRoom,
		RoomName:    roomTimeslotReq.RoomName,
		TimeSlotIDs: roomTimeslotReq.TimeSlotIDs,
	}

	roomTimeslot = service.RoomTimeslotRepository.AddRoomTimeslot(context, tx, roomTimeslot)
	
	res := rest.RoomTimeslotRequest{
		IDRoom:      roomTimeslot.IDRoom,
		RoomName:    roomTimeslot.RoomName,
		TimeSlotIDs: roomTimeslot.TimeSlotIDs,
	}
	
	return res, err
}
