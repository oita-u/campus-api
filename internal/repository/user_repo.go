package repository

import (
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(u *model.User) error {
	_, err := db.Conn.Exec(`INSERT INTO users (id, email, password, name) VALUES ($1, $2, $3, $4)`, u.ID, u.Email, u.Password, u.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	var u model.User
	err := db.Conn.Get(&u, `SELECT * FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var u model.User
	err := db.Conn.Get(&u, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
