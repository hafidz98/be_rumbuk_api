package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type FloorRepo interface {
	Create(context context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor
	SelectById(context context.Context, tx *sql.Tx, floorId int) (domain.Floor, error)
	SelectAll(context context.Context, tx *sql.Tx) []domain.Floor
}

type FloorRepoImpl struct{}

func NewFloorRepo() FloorRepo {
	return &FloorRepoImpl{}
}

func (repo *FloorRepoImpl) Create(context context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	query := "INSERT INTO floor(floor_name, building_id) VALUES(?,?)"
	res, err := tx.ExecContext(context, query, floor.Name, floor.BuildingID)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	floor.ID = int(id)
	return floor
}

func (repo *FloorRepoImpl) SelectById(context context.Context, tx *sql.Tx, floorId int) (domain.Floor, error) {
	query := "SELECT id, floor_name, building_id, created_at, updated_at FROM floor WHERE id = ?"
	row, err := tx.QueryContext(context, query, floorId)
	helper.PanicIfError(err)
	defer row.Close()

	floor := domain.Floor{}
	if row.Next() {
		err := row.Scan(
			&floor.ID,
			&floor.Name,
			&floor.BuildingID,
			&floor.CreatedAt,
			&floor.UpdatedAt,
		)
		helper.PanicIfError(err)
		return floor, nil
	}

	return floor, errors.New("floor not found")
}

func (repo *FloorRepoImpl) SelectAll(context context.Context, tx *sql.Tx) []domain.Floor {
	query := "SELECT id, floor_name, building_id, created_at, updated_at FROM floor"
	rows, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var floors []domain.Floor
	for rows.Next() {
		floor := domain.Floor{}
		err := rows.Scan(
			&floor.ID,
			&floor.Name,
			&floor.BuildingID,
			&floor.CreatedAt,
			&floor.UpdatedAt,
		)
		helper.PanicIfError(err)
		floors = append(floors, floor)
	}
	return floors
}

//Gedung dapat memiliki lantai dan/atau ruangan atau tidak sama sekali
//Lantai harus memiliki Gedung dan dapat memiliki Ruangan atau tidak sama sekali
//Ruangan harus memiliki Lantai dan Gedung
