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
	SelectAllReservation(context context.Context, tx *sql.Tx) []domain.ReservationDetail
	SelectByReservationId(context context.Context, tx *sql.Tx, reservationId int) (domain.Reservation, error)
	SelectRoomByRTSId(context context.Context, tx *sql.Tx, roomTimeSlotId int) (rest.RoomData, error)
	SelectReservationByStudentId(context context.Context, tx *sql.Tx, studentId string) []domain.Reservation
	Create(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation
	UpdateStatus(context context.Context, tx *sql.Tx, reserve domain.Reservation) domain.Reservation
}

type ReserveRoomRepoImpl struct{}

func NewReserveRoomRepo() ReserveRoomRepo {
	return &ReserveRoomRepoImpl{}
}

func (repo *ReserveRoomRepoImpl) SelectAllReservation(context context.Context, tx *sql.Tx) []domain.ReservationDetail {
	query :=
		`
	SELECT
    	r.id AS reservation_id,
    	r.reservation_date,
    	r.activity,
    	r.status,
    	r.student_id,
    	st.name AS student_name,
    	t.start_time,
    	t.end_time,
    	rm.room_name,	
    	f.floor_name,
    	b.building_name
	FROM reservation_ts r
	JOIN student st ON r.student_id = st.student_id
	JOIN room_time_slot ts ON r.room_timeslot_id = ts.id
	JOIN time_slot t ON ts.time_slot_id = t.id
	JOIN room rm ON ts.room_id = rm.id
	JOIN floor f ON rm.floor_id = f.id
	JOIN building b ON f.building_id = b.id
	ORDER BY 
    	r.reservation_date DESC,
		r.id DESC;
	`

	row, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer row.Close()

	var reservation []domain.ReservationDetail
	for row.Next() {
		resData := domain.ReservationDetail{}
		err := row.Scan(
			&resData.ID,
			&resData.BookDate,
			&resData.Activity,
			&resData.Status,
			&resData.StudentID,
			&resData.StudentName,
			&resData.StartTime,
			&resData.EndTime,
			&resData.RoomName,
			&resData.FloorName,
			&resData.BuildingName,
		)
		helper.PanicIfError(err)
		reservation = append(reservation, resData)
	}

	return reservation
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

func (repo *ReserveRoomRepoImpl) SelectRoomByRTSId(context context.Context, tx *sql.Tx, roomTimeSlotId int) (rest.RoomData, error) {
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

	return roomData, errors.New("room data not found")
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

	return reservation
}

func (repo *ReserveRoomRepoImpl) SelectByReservationId(context context.Context, tx *sql.Tx, reservationId int) (domain.Reservation, error) {
	query :=
		`
		SELECT res.id, res.student_id, res.reservation_date, res.activity, res.status, res.room_timeslot_id
		FROM reservation_ts res
		JOIN student s ON res.student_id = s.student_id
		WHERE res.id = ?;
	`
	row, err := tx.QueryContext(context, query, reservationId)
	helper.PanicIfError(err)
	defer row.Close()

	resData := domain.Reservation{}
	if row.Next() {
		err := row.Scan(
			&resData.ID,
			&resData.StudentID,
			&resData.BookDate,
			&resData.Activity,
			&resData.Status,
			&resData.RoomTimeSlotID,
		)
		helper.PanicIfError(err)
		return resData, nil
	}

	return resData, errors.New("reservation not found")
}
