package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/hafidz98/be_rumbuk_api/model/domain"
)

type StudentRepo interface {
	Create(context context.Context, tx *sql.Tx, student domain.Students) domain.Students
	Update(context context.Context, tx *sql.Tx, student domain.Students) domain.Students
	Delete(context context.Context, tx *sql.Tx, student domain.Students)
	FetchAll(context context.Context, tx *sql.Tx) []domain.Students
	FetchBySId(ctx context.Context, tx *sql.Tx, studentID string) (domain.Students, error)
}

type StudentRepoImpl struct{}

func NewStudentRepo() StudentRepo {
	return &StudentRepoImpl{}
}

func (repo *StudentRepoImpl) Create(context context.Context, tx *sql.Tx, student domain.Students) domain.Students {
	query := "insert into student(student_id, name, gender, batch_year, major, faculty, phone_number, email, password) values(?,?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(context, query, student.StudentID, student.Name, student.Gender, student.BatchYear, student.Major, student.Faculty, student.PhoneNumber, student.Email, student.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	student.ID = int(id)
	return student
}

func (repo *StudentRepoImpl) Update(context context.Context, tx *sql.Tx, student domain.Students) domain.Students {
	query := "update student set name = ? where student_id = ?"
	_, err := tx.ExecContext(context, query, student.Name, student.StudentID)
	helper.PanicIfError(err)

	return student
}

func (repo *StudentRepoImpl) Delete(context context.Context, tx *sql.Tx, category domain.Students) {
	query := "delete from student where student_id = ?"
	_, err := tx.ExecContext(context, query, category.StudentID)
	helper.PanicIfError(err)
}

func (repo *StudentRepoImpl) FetchAll(context context.Context, tx *sql.Tx) []domain.Students {
	query := "select student_id, name, gender, batch_year, major, faculty, phone_number, email from student"
	rows, err := tx.QueryContext(context, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var students []domain.Students
	for rows.Next() {
		student := domain.Students{}
		err := rows.Scan(
			&student.StudentID,
			&student.Name,
			&student.Gender,
			&student.BatchYear,
			&student.Major,
			&student.Faculty,
			&student.PhoneNumber,
			&student.Email,
		)
		helper.PanicIfError(err)
		students = append(students, student)
	}

	return students
}

func (repo *StudentRepoImpl) FetchBySId(context context.Context, tx *sql.Tx, studentId string) (domain.Students, error) {
	query := "select student_id, name, gender, batch_year, major, faculty, phone_number, email from student where student_id=?"
	rows, err := tx.QueryContext(context, query, studentId)
	helper.PanicIfError(err)
	defer rows.Close()

	student := domain.Students{}
	if rows.Next() {
		err := rows.Scan(
			&student.StudentID,
			&student.Name,
			&student.Gender,
			&student.BatchYear,
			&student.Major,
			&student.Faculty,
			&student.PhoneNumber,
			&student.Email,
		)
		helper.PanicIfError(err)
		return student, nil
	}

	return student, errors.New("student not found")
}
