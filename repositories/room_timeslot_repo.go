package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type RoomTimeslotRepo interface {
	AddRoomTimeslot(context context.Context, tx *sql.Tx, roomTimeslot domain.RoomTimeslot) domain.RoomTimeslot
}

type RoomTimeslotRepoImpl struct{}

func NewRoomTimeslotRepo() RoomTimeslotRepo {
	return &RoomTimeslotRepoImpl{}
}

func (repo *RoomTimeslotRepoImpl) AddRoomTimeslot(context context.Context, tx *sql.Tx, roomTimeslot domain.RoomTimeslot) domain.RoomTimeslot {
	// query := "DELETE FROM room_time_slot WHERE room_id = ?"
	// _, err := tx.ExecContext(context, query, roomTimeslot.IDRoom)
	// helper.PanicIfError(err)

	if len(roomTimeslot.TimeSlotIDs) > 0 {
		valueStrings := make([]string, 0, len(roomTimeslot.TimeSlotIDs))
		valueArgs := make([]interface{}, 0, len(roomTimeslot.TimeSlotIDs)*2)
		for _, tsID := range roomTimeslot.TimeSlotIDs {
			valueStrings = append(valueStrings, "(?, ?)")
			valueArgs = append(valueArgs, roomTimeslot.IDRoom, tsID)
		}

		stmt := fmt.Sprintf("INSERT INTO room_time_slot (room_id, time_slot_id) VALUES %s",
			strings.Join(valueStrings, ","))

		_, err := tx.ExecContext(context, stmt, valueArgs...)
		helper.PanicIfError(err)
	}

	return roomTimeslot
}
