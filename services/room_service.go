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

type RoomService interface {
	Create(context context.Context, request rest.RoomCreateRequest) rest.RoomResponse
	Update(context context.Context, request rest.RoomUpdateRequest) rest.RoomResponse
	UpdateRoomStatus(context context.Context, request rest.RoomUpdateRequest) rest.RoomResponse
	Delete(context context.Context, roomId int)
	FetchAll(context context.Context) []rest.RoomResponse
	FetchByID(context context.Context, roomId int) rest.RoomResponse
}

type RoomServiceImpl struct {
	RoomRepository repositories.RoomRepo
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoomService(roomRepository repositories.RoomRepo, DB *sql.DB, validate *validator.Validate) RoomService {
	return &RoomServiceImpl{
		RoomRepository: roomRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func toRoomResponse(room domain.Room) rest.RoomResponse {
	return rest.RoomResponse{
		ID:        room.ID,
		Name:      room.Name,
		Capacity:  room.Capacity,
		Building:  room.BuildingID,
		Floor:     room.FloorID,
		Status:    room.Status,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
}

func (service *RoomServiceImpl) Create(context context.Context, request rest.RoomCreateRequest) rest.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room := domain.Room{
		Name:       request.Name,
		Capacity:   request.Capacity,
		BuildingID: request.Building,
		FloorID:    request.Floor,
	}

	room = service.RoomRepository.Create(context, tx, room)

	return toRoomResponse(room)
}

func (service *RoomServiceImpl) Update(context context.Context, request rest.RoomUpdateRequest) rest.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room, err := service.RoomRepository.FetchByRoomID(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	room = domain.Room{
		ID:       request.ID,
		Name:     request.Name,
		Capacity: request.Capacity,
	}

	room = service.RoomRepository.Update(context, tx, room)

	return toRoomResponse(room)
}

func (service *RoomServiceImpl) UpdateRoomStatus(context context.Context, request rest.RoomUpdateRequest) rest.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room, err := service.RoomRepository.FetchByRoomID(context, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	room = domain.Room{
		ID:     request.ID,
		Status: request.Status,
	}

	room = service.RoomRepository.UpdateRoomStatus(context, tx, room)

	return toRoomResponse(room)
}

func (service *RoomServiceImpl) Delete(context context.Context, roomId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room, err := service.RoomRepository.FetchByRoomID(context, tx, roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoomRepository.Delete(context, tx, room)
}

func (service *RoomServiceImpl) FetchAll(context context.Context) []rest.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FetchAll(context, tx)

	var roomResponses []rest.RoomResponse
	for _, room := range rooms {
		room := rest.RoomResponse{
			ID:        room.ID,
			Name:      room.Name,
			Capacity:  room.Capacity,
			Building:  room.BuildingID,
			Floor:     room.FloorID,
			Status:    room.Status,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		}
		roomResponses = append(roomResponses, room)
	}

	return roomResponses
}

func (service *RoomServiceImpl) FetchByID(context context.Context, roomId int) rest.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	room, err := service.RoomRepository.FetchByRoomID(context, tx, roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return toRoomResponse(room)
}
