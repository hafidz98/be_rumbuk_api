package repositories

import (
	"context"
	"database/sql"
	"errors"

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

	staff := domain.Staff{
		StaffID:  staffID,
		Password: "staff7@pwd",
	}

	if staffID != "000007" {
		return staff, errors.New("staff not found")
	}

	return staff, nil
}
