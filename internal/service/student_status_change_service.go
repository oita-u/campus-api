package service

import (
	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type StudentStatusChangeService struct {
	Repo *repository.StudentStatusChangeRepository
}

func (s *StudentStatusChangeService) GetHistory(studentNumber string) ([]model.StudentStatusChange, error) {
	return s.Repo.GetByStudentNumber(studentNumber)
}
