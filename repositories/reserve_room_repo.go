package repositories

import (
	"context"
	"database/sql"
	"errors"
	//"fmt"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
	"github.com/hafidz98/be_rumbuk_api/models/rest"
)

type ReserveRoomRepo interface {
	//SelectAllReserveRoom()
	//SelectReserveRoomById()
	SelectReservationById(context context.Context, tx *sql.Tx, roomTimeSlotId int) (rest.RoomData, error)
	SelectReservationByStudentId(context context.Context, tx *sql.Tx, studentId string) []domain.Reservation
	Create(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation
	UpdateStatus(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation
}

type ReserveRoomRepoImpl struct{}

func NewReserveRoomRepo() ReserveRoomRepo {
	return &ReserveRoomRepoImpl{}
}

func (repo *ReserveRoomRepoImpl) Create(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation {
	query :=
		`
		INSERT INTO reservation_ts
		(
			reservation_date,
			activity,
			room_timeslot_id,
			student_id
		)
		VALUES(?, ?, ?, ?)
	`
	result, err := tx.ExecContext(context, query, reserve.BookDate, reserve.Activity, reserve.RoomTimeSlotID, reserve.StudentID)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	reserve.ID = int(id)
	return reserve
}

func (repo *ReserveRoomRepoImpl) UpdateStatus(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation {
	query :=
		`
		UPDATE reservation_ts
		SET status = ?
		WHERE id = ?
	`

	_, err := tx.ExecContext(context, query, reserve.Status, reserve.ID)
	helper.PanicIfError(err)

	return reserve
}

func (repo *ReserveRoomRepoImpl) SelectReservationById(context context.Context, tx *sql.Tx, roomTimeSlotId int) (rest.RoomData, error) {
	query :=
		`
		SELECT
		r.room_name, 
		f.floor_name, 
		b.building_name, 
		r.capacity, 
		ts.start_time,
		ts.end_time
		FROM room_time_slot rts
		JOIN time_slot ts ON ts.id = rts.time_slot_id
		JOIN room r ON r.id = rts.room_id
		JOIN floor f ON f.id = r.floor_id
		JOIN building b ON b.id = r.building_id
		WHERE rts.id = ?
	`
	row, err := tx.QueryContext(context, query, roomTimeSlotId)
	helper.PanicIfError(err)
	defer row.Close()

	roomData := rest.RoomData{}
	if row.Next() {
		err := row.Scan(
			&roomData.RoomName,
			&roomData.FloorName,
			&roomData.BuildingName,
			&roomData.Capacity,
			&roomData.StartTime,
			&roomData.EndTime,
		)
		helper.PanicIfError(err)
		return roomData, nil
	}

	return roomData, errors.New("not found")
}

func (repo *ReserveRoomRepoImpl) SelectReservationByStudentId(context context.Context, tx *sql.Tx, studentId string) []domain.Reservation {
	query :=
		`
		SELECT res.id, res.student_id, res.reservation_date, res.activity, res.status, res.room_timeslot_id
		FROM reservation_ts res
		JOIN student s ON res.student_id = s.student_id
		WHERE s.student_id = ?;
	`
	row, err := tx.QueryContext(context, query, studentId)
	helper.PanicIfError(err)
	defer row.Close()

	var reservation []domain.Reservation
	for row.Next() {
		resData := domain.Reservation{}
		err := row.Scan(
			&resData.ID,
			&resData.StudentID,
			&resData.BookDate,
			&resData.Activity,
			&resData.Status,
			&resData.RoomTimeSlotID,
		)
		helper.PanicIfError(err)
		reservation = append(reservation, resData)
	}

	//fmt.Printf("Data Repo: %v\n", reservation)

	return reservation
}
