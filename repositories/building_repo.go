package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type BuildingRepo interface {
	Create(context context.Context, tx *sql.Tx, building domain.Building) domain.Building
	Update(context context.Context, tx *sql.Tx, building domain.Building) domain.Building
	Delete(context context.Context, tx *sql.Tx, building domain.Building)
	FetchAll(context context.Context, tx *sql.Tx) []domain.Building
	FetchByID(ctx context.Context, tx *sql.Tx, buildingID int) (domain.Building, error)
}

type BuildingRepoImpl struct{}

func NewBuildingRepo() BuildingRepo {
	return &BuildingRepoImpl{}
}

func (repo *BuildingRepoImpl) Create(context context.Context, tx *sql.Tx, building domain.Building) domain.Building {
	query := "insert into building(building_name) values(?)"
	result, err := tx.ExecContext(context, query, building.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	building.ID = int(id)
	return building
}

func (repo *BuildingRepoImpl) Update(context context.Context, tx *sql.Tx, building domain.Building) domain.Building {
	query := "update building set building_name = ? where id = ?"
	_, err := tx.ExecContext(context, query, building.Name, building.ID)
	helper.PanicIfError(err)

	return building
}

func (repo *BuildingRepoImpl) Delete(context context.Context, tx *sql.Tx, building domain.Building) {}

func (repo *BuildingRepoImpl) FetchAll(context context.Context, tx *sql.Tx) []domain.Building {
	query := "select id, building_name, created_at, updated_at from building"
	rows, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var buildings []domain.Building
	for rows.Next() {
		building := domain.Building{}
		err := rows.Scan(
			&building.ID,
			&building.Name,
			&building.CreatedAt,
			&building.UpdatedAt,
		)
		helper.PanicIfError(err)
		buildings = append(buildings, building)
	}

	return buildings
}

func (repo *BuildingRepoImpl) FetchByID(context context.Context, tx *sql.Tx, buildingID int) (domain.Building, error) {
	query := "select id, building_name, created_at, updated_at from building where id=?"
	rows, err := tx.QueryContext(context, query, buildingID)
	helper.PanicIfError(err)
	defer rows.Close()

	building := domain.Building{}
	for rows.Next() {
		err := rows.Scan(
			&building.ID,
			&building.Name,
			&building.CreatedAt,
			&building.UpdatedAt,
		)
		helper.PanicIfError(err)
		return building, nil
	}

	return building, errors.New("building not found")
}
