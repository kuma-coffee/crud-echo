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

func (su StudentUsecase) GetStudent(id int) (domain.Student, error) {
	return su.StudentRepository.GetStudent(id)
}

func (su StudentUsecase) PostStudent(student domain.Student) error {
	return su.StudentRepository.PostStudent(student)
}

func (su StudentUsecase) UpdateStudent(id int, student domain.Student) error {
	return su.StudentRepository.UpdateStudent(id, student)
}

func (su StudentUsecase) DeleteStudent(id int) error {
	return su.StudentRepository.DeleteStudent(id)
}
