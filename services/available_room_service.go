package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type AvailableRoomService interface {
	GetAllAvailableRoom(context context.Context, params string) []rest.AvailabeRoomResponse
	GetAvailableRoom(context context.Context, date string, roomTimeslotId int) bool
}

type AvailableRoomServiceImpl struct {
	AvailableRoomRepo repositories.AvailableRoomRepo
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAvailableRoomService(roomRepository repositories.AvailableRoomRepo, DB *sql.DB, validate *validator.Validate) AvailableRoomService {
	return &AvailableRoomServiceImpl{
		AvailableRoomRepo: roomRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AvailableRoomServiceImpl) GetAllAvailableRoom(context context.Context, params string) []rest.AvailabeRoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	rooms := service.AvailableRoomRepo.SelectAllAvailableRoom(context, tx, params)

	//helper.Info.Printf("Rooms Data: %v", rooms)

	availableRoom := make(map[int]rest.AvailabeRoomResponse)
	for _, data := range rooms {

		//helper.Info.Printf("Room Data: %v", data)

		building, ok := availableRoom[data.Building.ID]

		if !ok {
			building = rest.AvailabeRoomResponse{
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
				ID:    data.Floor.ID,
				Name:  data.Floor.Name,
				Rooms: []rest.RoomResponse{},
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
			r := rest.RoomResponse{
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
			Duration:  data.TimeSlot.Duration,
			Reserved:  data.Reserved,
			RoomTSID:  data.RoomTimeSlotID,
		}

		//helper.Info.Printf("Room Data: %v", data.Reserved)

		building.Floors[floorIdx].Rooms[roomsIdx].TimeSlot = append(building.Floors[floorIdx].Rooms[roomsIdx].TimeSlot, ts)

		availableRoom[data.Building.ID] = building
	}

	availableRoomList := make([]rest.AvailabeRoomResponse, 0, len(availableRoom))
	for _, b := range availableRoom {
		availableRoomList = append(availableRoomList, b)
		//helper.Info.Printf("Room Data: %v", b)
	}

	return availableRoomList
}

func (service *AvailableRoomServiceImpl) GetAvailableRoom(context context.Context, date string, roomTimeslotId int) bool {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ok := service.AvailableRoomRepo.SelectIsReserveRoom(context, tx, date, roomTimeslotId)
	return ok
}
