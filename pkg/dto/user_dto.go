package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserDTO struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserDTO) ValidationUser() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Fullname, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Address, validation.Required),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 20)),
	)
}

func (l LoginDTO) ValidationLogin() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 20)),
	)
}
