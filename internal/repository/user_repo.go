package repository

import (
	"errors"

	"github.com/oita-u/campus-api/internal/model"
)

type UserRepository struct {
	data map[string]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: make(map[string]*model.User),
	}
}

func (r *UserRepository) Create(u *model.User) error {
	if _, ok := r.data[u.ID]; ok {
		return errors.New("user already exists")
	}
	r.data[u.ID] = u
	return nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {
	u, ok := r.data[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	for _, u := range r.data {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("not found")
}
