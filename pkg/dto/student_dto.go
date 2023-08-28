package dto

import validation "github.com/go-ozzo/ozzo-validation"

type StudentDTO struct {
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

func (s StudentDTO) ValidationStudent() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Fullname, validation.Required),
		validation.Field(&s.Birthdate, validation.Required),
		validation.Field(&s.Class, validation.Required),
		validation.Field(&s.Batch, validation.Required),
		validation.Field(&s.SchoolName, validation.Required),
	)
}
