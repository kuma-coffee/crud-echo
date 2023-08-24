package domain

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
	PostStudent(student Student) error
	UpdateStudent(id int, student Student) error
	DeleteStudent(id int) error
}

type StudentUsecase interface {
	GetStudents() ([]Student, error)
	PostStudent(student Student) error
	UpdateStudent(id int, student Student) error
	DeleteStudent(id int) error
}
