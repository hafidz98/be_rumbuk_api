package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type RoomRepo interface {
	Create(context context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Update(context context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Delete(context context.Context, tx *sql.Tx, room domain.Room)
	FetchAll(context context.Context, tx *sql.Tx) []domain.Room
	FetchAllRoomSpecial(context context.Context, tx *sql.Tx, params string) []domain.Rooms
	FetchByRoomID(context context.Context, tx *sql.Tx, roomId int) (domain.Room, error)
	FetchAllTS(context context.Context, tx *sql.Tx) []domain.TimeSlot
}

type RoomRepoImpl struct{}

func NewRoomRepo() RoomRepo {
	return &RoomRepoImpl{}
}

func (repo *RoomRepoImpl) Create(context context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	stmt := "INSERT INTO room(room_name, capacity, building_id, floor_id) VALUES(?,?,?,?)"
	res, err := tx.ExecContext(context, stmt, room.Name, room.Capacity, room.BuildingID, room.FloorID)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	room.ID = int(id)
	return room
}

func (repo *RoomRepoImpl) Update(context context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	stmt := "UPDATE room SET room_name = ?, capacity = ? WHERE id = ?"
	_, err := tx.ExecContext(context, stmt, room.Name, room.Capacity, room.ID)
	helper.PanicIfError(err)

	return room
}

func (repo *RoomRepoImpl) Delete(context context.Context, tx *sql.Tx, room domain.Room) {
	stmt := "UPDATE room SET is_deleted = ? WHERE id = ?"
	_, err := tx.ExecContext(context, stmt, true, room.ID)
	helper.PanicIfError(err)
}

func (repo *RoomRepoImpl) FetchByRoomID(context context.Context, tx *sql.Tx, roomId int) (domain.Room, error) {
	stmt := "SELECT id, room_name, capacity, building_id, floor_id, created_at, updated_at FROM room WHERE id = ?"
	rows, err := tx.QueryContext(context, stmt, roomId)
	helper.PanicIfError(err)
	defer rows.Close()

	room := domain.Room{}
	if rows.Next() {
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Capacity,
			&room.BuildingID,
			&room.FloorID,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		helper.PanicIfError(err)
		return room, nil
	}

	return room, errors.New("room not found")
}

func (repo *RoomRepoImpl) FetchAll(context context.Context, tx *sql.Tx) []domain.Room {
	stmt := `
		SELECT
			r.id,
			r.room_name,
			r.capacity,
			b.id,
			f.id,
			r.created_at,
			r.updated_at
		FROM room r
			JOIN floor f ON r.floor_id = f.id
			JOIN building b on r.building_id = b.id;
	`
	rows, err := tx.QueryContext(context, stmt)
	helper.PanicIfError(err)
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		room := domain.Room{}
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Capacity,
			&room.BuildingID,
			&room.FloorID,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		helper.PanicIfError(err)
		rooms = append(rooms, room)
	}
	return rooms
}

func (repo *RoomRepoImpl) FetchAllRoomSpecial(context context.Context, tx *sql.Tx, params string) []domain.Rooms {
	stmt := `
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
			CASE
				WHEN EXISTS(
					SELECT 1
					FROM reservation_ts rsts
						JOIN room_time_slot rts ON rsts.room_timeslot_id = rts.id
					WHERE
						rts.room_id = r.id AND rts.time_slot_id = ts.id
						AND rsts.reservation_date = ?
				) THEN true
				ELSE false
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

	rows, err := tx.QueryContext(context, stmt, params)
	helper.PanicIfError(err)
	defer rows.Close()

	var rooms []domain.Rooms
	for rows.Next() {
		var b domain.Building
		var f domain.Floor
		var rr domain.Room
		var ts domain.TimeSlot
		var r domain.Rooms

		err := rows.Scan(
			&b.ID,
			&b.Name,
			&f.ID,
			&f.Name,
			&rr.ID,
			&rr.Name,
			&rr.Capacity,
			&ts.ID,
			&ts.StartTime,
			&ts.EndTime,
			&r.Reserved,
		)

		r.Building = b
		r.Floor = f
		r.Room = rr
		r.TimeSlot = ts

		helper.PanicIfError(err)
		rooms = append(rooms, r)
	}

	return rooms
}

func (repo *RoomRepoImpl) FetchAllTS(context context.Context, tx *sql.Tx) []domain.TimeSlot {
	stmt := "SELECT ts.id, ts.start_time, ts.end_time, ts.created_at FROM time_slot ts"
	rows, err := tx.QueryContext(context, stmt)
	helper.PanicIfError(err)
	defer rows.Close()

	var tss []domain.TimeSlot
	for rows.Next() {
		ts := domain.TimeSlot{}
		err := rows.Scan(
			&ts.ID,
			&ts.StartTime,
			&ts.EndTime,
			&ts.CreatedAt,
		)
		helper.PanicIfError(err)
		tss = append(tss, ts)
	}
	return tss
}
