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
	UpdateRoomStatus(context context.Context, tx *sql.Tx, status domain.Room) domain.Room
	Delete(context context.Context, tx *sql.Tx, room domain.Room)
	FetchAll(context context.Context, tx *sql.Tx) []domain.Room
	FetchByRoomID(context context.Context, tx *sql.Tx, roomId int) (domain.Room, error)
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

// Status yang terdapat pada Ruangan.
//
// 0 : Ruangan tidak dapat dilakukan proses peminjaman.
// 1 : (Default) Ruangan dapat dipinjam.
// 2 : Lain-lain (Dapat diatur kemudian).
func (repo *RoomRepoImpl) UpdateRoomStatus(context context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	stmt := "UPDATE room SET status = ? WHERE id = ?"
	_, err := tx.ExecContext(context, stmt, room.Status, room.ID)
	helper.PanicIfError(err)

	return room
}

func (repo *RoomRepoImpl) Delete(context context.Context, tx *sql.Tx, room domain.Room) {
	stmt := "UPDATE room SET is_deleted = ? WHERE id = ?"
	_, err := tx.ExecContext(context, stmt, true, room.ID)
	helper.PanicIfError(err)
}

func (repo *RoomRepoImpl) FetchByRoomID(context context.Context, tx *sql.Tx, roomId int) (domain.Room, error) {
	stmt := "SELECT id, room_name, capacity, building_id, floor_id, status, created_at, updated_at FROM room WHERE id = ?"
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
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		helper.PanicIfError(err)
		return room, nil
	}

	return room, errors.New("room not found")
}

func (repo *RoomRepoImpl) FetchAll(context context.Context, tx *sql.Tx) []domain.Room {
	stmt :=
		`
		SELECT
			r.id,
			r.room_name,
			r.capacity,
			b.id,
			f.id,
			r.status,
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
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		helper.PanicIfError(err)
		rooms = append(rooms, room)
	}
	return rooms
}
