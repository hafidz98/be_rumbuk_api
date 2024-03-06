package repositories

import (
	"context"
	"database/sql"

	//"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type AvailableRoomRepo interface {
	SelectIsReserveRoom(context context.Context, tx *sql.Tx, date string, roomTimeslotId int) bool
	SelectAllAvailableRoom(context context.Context, tx *sql.Tx, params string) []domain.AvailableRoom
}

type AvailableRoomRepoImpl struct{}

func NewAvailableRoomRepo() AvailableRoomRepo {
	return &AvailableRoomRepoImpl{}
}

func (repo *AvailableRoomRepoImpl) SelectAllAvailableRoom(context context.Context, tx *sql.Tx, params string) []domain.AvailableRoom {
	query := `
		SELECT
			b.id,
			b.building_name,
			f.id,
			f.floor_name,
			r.id,
			r.room_name,
			r.capacity,
			ts.id,
			ts.start_time,
			ts.end_time,
			ts.duration,
			rts.id,
			CASE
				WHEN EXISTS(
					SELECT 1
					FROM reservation_ts rsts
						JOIN room_time_slot rts ON rsts.room_timeslot_id = rts.id
					WHERE
						rts.room_id = r.id AND rts.time_slot_id = ts.id
						AND rsts.reservation_date = ?
				) THEN '1'
				ELSE '0'
			END AS 'reserved'
		FROM room r
			JOIN floor f ON r.floor_id = f.id
			JOIN building b ON b.id = f.building_id
			JOIN room_time_slot rts ON r.id = rts.room_id
			JOIN time_slot ts ON rts.time_slot_id = ts.id
		ORDER BY
			b.building_name asc,
			f.floor_name ASC,
			r.room_name ASC,
			ts.start_time ASC
	`

	rows, err := tx.QueryContext(context, query, params)
	helper.PanicIfError(err)
	defer rows.Close()

	var availableRoom []domain.AvailableRoom
	for rows.Next() {
		var building domain.Building
		var floor domain.Floor
		var room domain.Room
		var timeslot domain.TimeSlot
		var availableroom domain.AvailableRoom

		err := rows.Scan(
			&building.ID,
			&building.Name,
			&floor.ID,
			&floor.Name,
			&room.ID,
			&room.Name,
			&room.Capacity,
			&timeslot.ID,
			&timeslot.StartTime,
			&timeslot.EndTime,
			&timeslot.Duration,
			&availableroom.RoomTimeSlotID,
			&availableroom.Reserved,
		)

		availableroom.Building = building
		availableroom.Floor = floor
		availableroom.Room = room
		availableroom.TimeSlot = timeslot

		helper.PanicIfError(err)
		availableRoom = append(availableRoom, availableroom)
	}

	return availableRoom
}

func (repo *AvailableRoomRepoImpl) SelectIsReserveRoom(context context.Context, tx *sql.Tx, date string, roomTimeslotId int) bool {
	query := `
		SELECT
			CASE
				WHEN EXISTS(
					SELECT 1
					FROM reservation_ts rsts
						JOIN room_time_slot rts ON rsts.room_timeslot_id = rts.id
					WHERE
						rts.room_id = r.id AND rts.time_slot_id = ts.id
						AND rsts.reservation_date = ?
				) THEN '1'
				ELSE '0'
			END AS 'reserved'
		FROM room r
			JOIN floor f ON r.floor_id = f.id
			JOIN building b ON b.id = f.building_id
			JOIN room_time_slot rts ON r.id = rts.room_id
			JOIN time_slot ts ON rts.time_slot_id = ts.id
		WHERE rts.id = ?
		ORDER BY
			b.building_name asc,
			f.floor_name ASC,
			r.room_name ASC,
			ts.start_time ASC
	`

	rows, err := tx.QueryContext(context, query, date, roomTimeslotId)
	helper.PanicIfError(err)
	defer rows.Close()

	var availableroom domain.AvailableRoom
	if rows.Next() {
		err := rows.Scan(
			&availableroom.Reserved,
		)
		helper.PanicIfError(err)
	}

	return *availableroom.Reserved
}