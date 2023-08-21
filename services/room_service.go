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
	Delete(context context.Context, roomId int)
	FetchAll(context context.Context) []rest.RoomResponse
	FetchByID(context context.Context, roomId int) rest.RoomResponse
	FetchAllRooms(context context.Context, params string) []rest.BuildingResponse
	FetchAllTS(context context.Context) []rest.TimeSlotResponse
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
		ID:       room.ID,
		Name:     room.Name,
		Capacity: room.Capacity,
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

func (service *RoomServiceImpl) Delete(context context.Context, roomId int) {}

func (service *RoomServiceImpl) FetchAll(context context.Context) []rest.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FetchAll(context, tx)

	var roomResponses []rest.RoomResponse
	for _, room := range rooms {
		room := rest.RoomResponse{
			ID:       room.ID,
			Name:     room.Name,
			Capacity: room.Capacity,
			Building: room.BuildingID,
			Floor:    room.FloorID,
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

func (service *RoomServiceImpl) FetchAllRooms(context context.Context, params string) []rest.BuildingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FetchAllRoomSpecial(context, tx, params)

	buildings := make(map[int]rest.BuildingResponse)
	for _, data := range rooms {
		building, ok := buildings[data.Building.ID]

		if !ok {
			building = rest.BuildingResponse{
				ID:     data.Building.ID,
				Name:   data.Building.Name,
				Floors: []rest.FloorResponse{},
			}
		}

		floorIdx := -1
		for i, f := range building.Floors {
			if f.ID == data.Floor.ID {
				floorIdx = i
				break
			}
		}

		if floorIdx == -1 {
			f := rest.FloorResponse{
				ID:     data.Floor.ID,
				Number: data.Floor.Name,
				Rooms:  []rest.Rooms{},
			}
			building.Floors = append(building.Floors, f)
			floorIdx = len(building.Floors) - 1
		}

		roomsIdx := -1
		for i, r := range building.Floors[floorIdx].Rooms {
			if r.ID == data.Room.ID {
				roomsIdx = i
				break
			}
		}

		if roomsIdx == -1 {
			r := rest.Rooms{
				ID:       data.Room.ID,
				Name:     data.Room.Name,
				Capacity: data.Room.Capacity,
				TimeSlot: []rest.TimeSlotResponse{},
			}
			building.Floors[floorIdx].Rooms = append(building.Floors[floorIdx].Rooms, r)
			roomsIdx = len(building.Floors[floorIdx].Rooms) - 1
		}

		ts := rest.TimeSlotResponse{
			ID:        data.TimeSlot.ID,
			StartTime: data.TimeSlot.StartTime,
			EndTime:   data.TimeSlot.EndTime,
			Reserved:  data.Reserved,
		}

		building.Floors[floorIdx].Rooms[roomsIdx].TimeSlot = append(building.Floors[floorIdx].Rooms[roomsIdx].TimeSlot, ts)

		buildings[data.Building.ID] = building
	}

	buildingList := make([]rest.BuildingResponse, 0, len(buildings))
	for _, b := range buildings {
		buildingList = append(buildingList, b)
	}

	return buildingList
}

func (service *RoomServiceImpl) FetchAllTS(context context.Context) []rest.TimeSlotResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.RoomRepository.FetchAllTS(context, tx)
	var roomResponses []rest.TimeSlotResponse

	for _, r := range rooms {
		r := rest.TimeSlotResponse{
			StartTime: r.StartTime,
			EndTime:   r.EndTime,
		}

		roomResponses = append(roomResponses, r)
	}

	return roomResponses
}
