package service

import (
	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type TimetableService struct {
	repo *repository.TimetableRepository
}

func NewTimetableService(r *repository.TimetableRepository) *TimetableService {
	return &TimetableService{repo: r}
}

func (s *TimetableService) GetTimetable(userID string, year int, semester string) ([]model.TimetableItem, int, string, error) {
	year, semester = resolveYearSemester(year, semester)

	items, err := s.repo.GetTimetableByUserID(userID, year, semester)
	if err != nil {
		return nil, 0, "", err
	}

	return items, year, semester, nil
}
