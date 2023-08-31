package domain

import "github.com/kuma-coffee/crud-echo/pkg/dto"

type Student struct {
	Id         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

type StudentRepository interface {
	GetStudents() ([]Student, error)
	GetStudentById(id int) (Student, error)
	PostStudent(student Student) error
	UpdateStudent(id int, student Student) error
	DeleteStudent(id int) error
	SearchStudent(query []string) ([]Student, error)
	SortStudent(query string) ([]Student, error)
}

type StudentUsecase interface {
	GetStudents() ([]Student, error)
	GetStudentById(id int) (Student, error)
	PostStudent(studentDTO dto.StudentDTO) error
	UpdateStudent(id int, studentDTO dto.StudentDTO) error
	DeleteStudent(id int) error
	SearchStudent(query []string) ([]Student, error)
	SortStudent(query string) ([]Student, error)
}
