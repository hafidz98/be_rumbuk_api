package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type TimeSlotRepo interface {
	Create(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot) domain.TimeSlot
	Update(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot) domain.TimeSlot
	Delete(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot)
	SelectById(context context.Context, tx *sql.Tx, timeslotId int) (domain.TimeSlot, error)
	SelectAll(context context.Context, tx *sql.Tx) []domain.TimeSlot
}

type TimeSlotRepoImpl struct{}

func NewTimeSlotRepo() TimeSlotRepo {
	return &TimeSlotRepoImpl{}
}

func (repo *TimeSlotRepoImpl) Create(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot) domain.TimeSlot {
	stmt := "INSERT INTO time_slot(start_time, end_time, duration) VALUES(?,?,?)"
	res, err := tx.ExecContext(context, stmt, timeslot.StartTime, timeslot.EndTime, timeslot.Duration)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	timeslot.ID = int(id)
	return timeslot
}

func (repo *TimeSlotRepoImpl) Update(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot) domain.TimeSlot {
	stmt := "UPDATE time_slot SET start_time = ?, end_time = ?, duration = ? WHERE id = ?"
	_, err := tx.ExecContext(context, stmt, timeslot.StartTime, timeslot.EndTime, timeslot.Duration, timeslot.ID)
	helper.PanicIfError(err)

	return timeslot
}

func (repo *TimeSlotRepoImpl) Delete(context context.Context, tx *sql.Tx, timeslot domain.TimeSlot) {

}

func (repo *TimeSlotRepoImpl) SelectById(context context.Context, tx *sql.Tx, timeslotId int) (domain.TimeSlot, error) {
	stmt := "SELECT ts.id, ts.start_time, ts.end_time, ts.duration, ts.created_at, ts.updated_at FROM time_slot ts WHERE id = ?"
	rows, err := tx.QueryContext(context, stmt, timeslotId)
	helper.PanicIfError(err)
	defer rows.Close()

	timeslot := domain.TimeSlot{}
	if rows.Next() {
		err := rows.Scan(
			&timeslot.ID,
			&timeslot.StartTime,
			&timeslot.EndTime,
			&timeslot.Duration,
			&timeslot.CreatedAt,
			&timeslot.UpdatedAt,
		)
		helper.PanicIfError(err)
		return timeslot, nil
	}

	return timeslot, errors.New("timeslot not found")
}

func (repo *TimeSlotRepoImpl) SelectAll(context context.Context, tx *sql.Tx) []domain.TimeSlot {
	stmt := "SELECT ts.id, ts.start_time, ts.end_time, ts.duration, ts.created_at, ts.updated_at FROM time_slot ts"
	rows, err := tx.QueryContext(context, stmt)
	helper.PanicIfError(err)
	defer rows.Close()

	var timeslot []domain.TimeSlot
	for rows.Next() {
		ts := domain.TimeSlot{}
		err := rows.Scan(
			&ts.ID,
			&ts.StartTime,
			&ts.EndTime,
			&ts.Duration,
			&ts.CreatedAt,
			&ts.UpdatedAt,
		)
		helper.PanicIfError(err)
		timeslot = append(timeslot, ts)
	}
	return timeslot
}
