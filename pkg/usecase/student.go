package usecase

import "github.com/kuma-coffee/crud-echo/pkg/domain"

type StudentUsecase struct {
	StudentRepository domain.StudentRepository
}

func NewStudentUsecase(studentRepository domain.StudentRepository) domain.StudentUsecase {
	return &StudentUsecase{studentRepository}
}

func (su StudentUsecase) GetStudents() ([]domain.Student, error) {
	return su.StudentRepository.GetStudents()
}
