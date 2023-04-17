package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/models/domain"
)

type StaffRepo interface {
	Create(context context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff
	Update(context context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff
	SoftDelete(context context.Context, tx *sql.Tx, staff domain.Staff)
	Delete(context context.Context, tx *sql.Tx, staff domain.Staff)
	CountAll(context context.Context, tx *sql.Tx) int
	FetchAll(context context.Context, tx *sql.Tx) []domain.Staff
	FetchAllFilter(context context.Context, tx *sql.Tx, filter *domain.FilterParams) []domain.Staff
	FetchById(context context.Context, tx *sql.Tx, staffID string) (domain.Staff, error)
}

type StaffRepoImpl struct{}

func NewStaffRepo() StaffRepo {
	return &StaffRepoImpl{}
}

func queryRowsWithFilterBuilder(query string, filter *domain.FilterParams) (q string, f []interface{}, err error) {
	var filterValues []interface{}

	offset := (filter.Page - 1) * filter.PerPage
	// pagination
	// if offset > 0 {
	// 	offset = offset - 1
	// }
	filterValues = append(filterValues, filter.PerPage)
	//query += "LIMIT ?" + strconv.Itoa(len(filterValues))
	query += ` LIMIT ? `
	filterValues = append(filterValues, offset)
	//query += "OFFSET ?" + strconv.Itoa(len(filterValues))
	query += ` OFFSET ? `

	return query, filterValues, err
}

func (repo *StaffRepoImpl) Create(context context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff {
	stmt := "INSERT INTO staff(staff_id, name, role, status, email, password) VALUES(?,?,?,?,?,?)"
	result, err := tx.ExecContext(context, stmt, staff.StaffID, staff.Name, staff.Role, staff.Status, staff.Email, staff.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	staff.ID = int(id)
	return staff
}

func (repo *StaffRepoImpl) Update(context context.Context, tx *sql.Tx, staff domain.Staff) domain.Staff {
	stmt := "UPDATE staff SET name = ?, email = ?, status = ? WHERE staff_id = ?"
	_, err := tx.ExecContext(context, stmt, staff.Name, staff.Email, staff.Status, staff.StaffID)
	helper.PanicIfError(err)
	return staff
}

func (repo *StaffRepoImpl) SoftDelete(context context.Context, tx *sql.Tx, staff domain.Staff) {
	stmt := "UPDATE staff SET status = ? WHERE staff_id = ?"
	_, err := tx.ExecContext(context, stmt, staff.Status, staff.StaffID)
	helper.PanicIfError(err)
}

func (repo *StaffRepoImpl) Delete(context context.Context, tx *sql.Tx, staff domain.Staff) {
	stmt := "DELETE FROM staff WHERE staff_id = ?"
	_, err := tx.ExecContext(context, stmt, staff.StaffID)
	helper.PanicIfError(err)
}

func (repo *StaffRepoImpl) FetchAll(context context.Context, tx *sql.Tx) []domain.Staff {
	stmt := "SELECT id, staff_id, name, role, status, email FROM staff WHERE status = 1"
	rows, err := tx.QueryContext(context, stmt)
	helper.PanicIfError(err)
	defer rows.Close()

	var staffs []domain.Staff
	for rows.Next() {
		staff := domain.Staff{}
		err := rows.Scan(
			&staff.ID,
			&staff.StaffID,
			&staff.Name,
			&staff.Role,
			&staff.Status,
			&staff.Email,
		)
		helper.PanicIfError(err)
		staffs = append(staffs, staff)
	}

	return staffs
}

func (repo *StaffRepoImpl) FetchAllFilter(context context.Context, tx *sql.Tx, filter *domain.FilterParams) []domain.Staff {
	stmt := "SELECT id, staff_id, name, role, status, email FROM staff WHERE status = 1"

	q, f, err := queryRowsWithFilterBuilder(stmt, filter)
	helper.PanicIfError(err)
	helper.Info.Printf("Query: %s", q)
	helper.Info.Printf("Params: %v", f)
	helper.Info.Printf("Params provided: %v", filter)

	rows, err := tx.QueryContext(context, q, f...)
	helper.PanicIfError(err)
	defer rows.Close()

	var staffs []domain.Staff
	for rows.Next() {
		staff := domain.Staff{}
		err := rows.Scan(
			&staff.ID,
			&staff.StaffID,
			&staff.Name,
			&staff.Role,
			&staff.Status,
			&staff.Email,
		)
		helper.PanicIfError(err)
		staffs = append(staffs, staff)
	}

	return staffs
}

func (repo *StaffRepoImpl) FetchById(context context.Context, tx *sql.Tx, staffID string) (domain.Staff, error) {
	query := "SELECT staff_id, name, role, email, password FROM staff WHERE staff_id=? & status = 1"
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

	return staff, errors.New("staff not found")
}

func (repo *StaffRepoImpl) CountAll(context context.Context, tx *sql.Tx) int {
	var count int
	stmt := `SELECT COUNT(*) FROM staff`
	err := tx.QueryRowContext(context, stmt).Scan(&count)
	helper.PanicIfError(err)
	return count
}
