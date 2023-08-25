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

func (sr StudentRepository) GetStudent(id int) (domain.Student, error) {
	student := domain.Student{}

	sql := `SELECT * FROM students WHERE id = $1`

	err := sr.db.QueryRow(sql, id).Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)

	return student, err
}

func (sr StudentRepository) PostStudent(student domain.Student) error {
	sql := `INSERT INTO students (fullname, address, birthdate, class, batch, school_name) VALUES($1, $2, $3, $4, $5, $6)`

	_, err := sr.db.Exec(sql, student.Fullname, student.Address, student.Birthdate, student.Class, student.Batch, student.SchoolName)

	return err
}

func (sr StudentRepository) UpdateStudent(id int, student domain.Student) error {
	sql := `UPDATE students SET fullname = $2, address = $3, birthdate = $4, class = $5, batch = $6, school_name = $7 WHERE id = $1`

	_, err := sr.db.Exec(sql, id, student.Fullname, student.Address, student.Birthdate, student.Class, student.Batch, student.SchoolName)

	return err
}

func (sr StudentRepository) DeleteStudent(id int) error {
	sql := `DELETE FROM students WHERE id = $1`

	_, err := sr.db.Exec(sql, id)

	return err
}
