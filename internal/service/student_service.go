package service

import (
	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func (s *StudentService) GetByID(id int) (*model.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) GetByStudentNumber(num string) (*model.Student, error) {
	return s.Repo.GetByStudentNumber(num)
}

func (s *StudentService) List() ([]model.Student, error) {
	return s.Repo.List()
}
