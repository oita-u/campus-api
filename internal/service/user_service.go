package service

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(email, password, name string) (*model.User, error) {
	_, err := s.repo.GetByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{
		ID:       uuid.New().String(),
		Email:    email,
		Password: string(hash),
		Name:     name,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Get(id string) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Login(email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("メールアドレスまたはパスワードが正しくありません")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("メールアドレスまたはパスワードが正しくありません")
	}

	return user, nil
}