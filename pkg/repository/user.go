package repository

import (
	"database/sql"

	"github.com/kuma-coffee/crud-echo/pkg/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{db}
}

func (ur UserRepository) Register(user domain.User) error {
	sql := `INSERT INTO users(fullname, email, address, password) VALUES($1, $2, $3, $4)`

	_, err := ur.db.Exec(sql, user.Fullname, user.Email, user.Address, user.Password)

	return err
}

func (ur UserRepository) GetUserEmail(email string) (domain.User, error) {
	user := domain.User{}

	sql := `SELECT * FROM users WHERE email = $1`

	err := ur.db.QueryRow(sql, email).Scan(&user.Id, &user.Fullname, &user.Email, &user.Address, &user.Password)
	return user, err
}
