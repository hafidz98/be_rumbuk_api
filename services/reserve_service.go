package services

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hafidz98/be_rumbuk_api/exception"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
	"github.com/hafidz98/be_rumbuk_api/repositories"
)

type ReservationService interface {
	SelectReservationByStudentID(context context.Context, studentId string) []rest.ReserveResponse
	CreateReservation(context context.Context, request rest.ReserveCreateRequest) (rest.ReserveResponse, string)
	CancelReservation(context context.Context, reservationId int)
	GetAllReservation(context context.Context) []rest.ReservationDetailResponse
}

type ReservationServiceImpl struct {
	ReserveRepo repositories.ReserveRoomRepo
	DB          *sql.DB
	Validate    *validator.Validate
}

func NewReservationService(reserveRepo repositories.ReserveRoomRepo, DB *sql.DB, validate *validator.Validate) ReservationService {
	return &ReservationServiceImpl{
		ReserveRepo: reserveRepo,
		DB:          DB,
		Validate:    validate,
	}
}

func toReserveResponse(reserve domain.Reservation, room rest.RoomData, statusText string) rest.ReserveResponse {
	return rest.ReserveResponse{
		ReserveID:  reserve.ID,
		StudentID:  reserve.StudentID,
		Activity:   reserve.Activity,
		Status:     reserve.Status,
		BookDate:   reserve.BookDate,
		Room:       &room,
		StatusText: statusText,
	}
}

func (service *ReservationServiceImpl) GetAllReservation(context context.Context) []rest.ReservationDetailResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	reservations := service.ReserveRepo.SelectAllReservation(context, tx)

	var reserveResponses []rest.ReservationDetailResponse
	for _, reservation := range reservations {
		reserveResponses = append(reserveResponses, rest.ReservationDetailResponse{
			ID:           reservation.ID,
			BookDate:     reservation.BookDate,
			StudentID:    reservation.StudentID,
			Activity:     reservation.Activity,
			Status:       reservation.Status,
			StudentName:  reservation.StudentName,
			StartTime:    reservation.StartTime,
			EndTime:      reservation.EndTime,
			RoomName:     reservation.RoomName,
			FloorName:    reservation.FloorName,
			BuildingName: reservation.BuildingName,
		})
	}

	return reserveResponses
}

var (
	bookings = make(map[string]time.Time)
	mu       sync.Mutex
)

func (service *ReservationServiceImpl) CreateReservation(context context.Context, request rest.ReserveCreateRequest) (rest.ReserveResponse, string) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	key := fmt.Sprintf("%s|%d", request.BookDate, request.RoomTimeSlotID)
	createdAt := time.Now()

	mu.Lock()
	defer mu.Unlock()

	formatedDate, err := time.Parse("2006-01-02", request.BookDate)
	helper.PanicIfError(err)

	if _, exists := bookings[key]; exists {
		return rest.ReserveResponse{
			ReserveID: 0,
			BookDate:  formatedDate,
			StudentID: request.StudentID,
			Activity:  request.Activity,
			Status:    "",
			Room:      nil,
		}, "already_booked"
	}

	// if existing, exists := bookings[key]; exists {
	// 	return rest.ReserveResponse{
	// 		ReserveID:  0,
	// 		BookDate:   formatedDate,
	// 		StudentID:  request.StudentID,
	// 		Activity:   request.Activity,
	// 		Status:     "",
	// 		Room:       nil,
	// 	}, fmt.Sprintf("already_booked_at_%s", existing.Format(time.RFC3339))
	// }

	bookings[key] = createdAt

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	formatedDate, err = time.Parse("2006-01-02", request.BookDate)
	helper.PanicIfError(err)

	reserve := domain.Reservation{
		StudentID:      request.StudentID,
		Activity:       request.Activity,
		BookDate:       formatedDate,
		RoomTimeSlotID: request.RoomTimeSlotID,
	}

	check := repositories.NewAvailableRoomRepo().SelectIsReserveRoom(context, tx, request.BookDate, request.RoomTimeSlotID)
	if check {
		return rest.ReserveResponse{}, "alreadey_reserved"
	}

	reserve = service.ReserveRepo.Create(context, tx, reserve)

	roomData, err := service.ReserveRepo.SelectRoomByRTSId(context, tx, reserve.RoomTimeSlotID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	reserve.Status = "1"

	return toReserveResponse(reserve, roomData, statusText(reserve.Status)), "success_booked"
}

func (service *ReservationServiceImpl) SelectReservationByStudentID(context context.Context, studentId string) []rest.ReserveResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	localDate := time.Now().Local()
	localTime := time.Now().Format("15:04")

	reservation := service.ReserveRepo.SelectReservationByStudentId(context, tx, studentId)

	var resData []rest.ReserveResponse
	for _, res := range reservation {
		reservationDate := res.BookDate
		daysDifference := math.Ceil(reservationDate.Sub(localDate).Hours() / 24)
		fmt.Printf("Tanggal Lokal: %v, Tanggal Reservasi: %v, Selisih hari: %v\n\n", localDate, reservationDate, daysDifference)

		roomData, err := service.ReserveRepo.SelectRoomByRTSId(context, tx, res.RoomTimeSlotID)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		if dayEqualZero(daysDifference) {
			if timeInBetwen(localTime, roomData.StartTime, roomData.EndTime) {
				res.Status = "2"
				service.ReserveRepo.UpdateStatus(context, tx, res)
			} else if localTime >= roomData.EndTime {
				res.Status = "3"
				service.ReserveRepo.UpdateStatus(context, tx, res)
			}
		} else if dayLowerThanZero(daysDifference) {
			if res.Status != "0" {
				res.Status = "3"
				service.ReserveRepo.UpdateStatus(context, tx, res)
			}
		}

		resData = append(resData, toReserveResponse(res, roomData, statusText(res.Status)))
	}

	return resData
}

func (service *ReservationServiceImpl) CancelReservation(context context.Context, reservationId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	reservation, err := service.ReserveRepo.SelectByReservationId(context, tx, reservationId)
	helper.PanicIfError(err)

	reservation = domain.Reservation{
		ID:     reservationId,
		Status: "0",
	}

	// roomData, err := service.ReserveRepo.SelectRoomByRTSId(context, tx, reservation.RoomTimeSlotID)
	// helper.PanicIfError(err)

	service.ReserveRepo.UpdateStatus(context, tx, reservation)
}

func statusText(statusCode string) string {
	switch statusCode {
	case "0":
		return "Canceled"
	case "1":
		return "Upcoming"
	case "2":
		return "Ongoing"
	case "3":
		return "Done"
	default:
		return "Invalid"
	}
}

func dayEqualZero(daysDifference float64) bool {
	return daysDifference == 0
}

// func dayGreaterThanZero(daysDifference float64) bool {
// 	return daysDifference >= 0
// }

func dayLowerThanZero(daysDifference float64) bool {
	return daysDifference <= 0
}

func timeInBetwen(localTime, startTime, endTime string) bool {
	return localTime >= startTime && localTime <= endTime
}
