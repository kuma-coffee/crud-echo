package repository

import (
	"database/sql"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) domain.StudentRepository {
	return &StudentRepository{db}
}

func (sr StudentRepository) GetStudents() ([]domain.Student, error) {
	students := []domain.Student{}

	sql := `SELECT * FROM students`

	rows, err := sr.db.Query(sql)
	for rows.Next() {
		student := domain.Student{}

		err := rows.Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, err
}
