package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type StaffRepo interface {
	FetchById(context context.Context, tx *sql.Tx, staffID string) (domain.Staff, error)
}

type StaffRepoImpl struct{}

func NewStaffRepo() StaffRepo {
	return &StaffRepoImpl{}
}

func (repo *StaffRepoImpl) FetchById(context context.Context, tx *sql.Tx, staffID string) (domain.Staff, error) {
	query := "select staff_id, name, role, email, password from staff where staff_id=?"
	rows, err := tx.QueryContext(context, query, staffID)
	helper.PanicIfError(err)
	defer rows.Close()

	staff := domain.Staff{}
	if rows.Next() {
		err := rows.Scan(
			&staff.StaffID,
			&staff.Name,
			&staff.Role,
			&staff.Email,
			&staff.Password,
		)
		helper.PanicIfError(err)
		return staff, nil
	}

	return staff, errors.New("Staff not found")
}
