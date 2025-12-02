package service

import (
	"time"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type TimetableService struct {
	repo *repository.TimetableRepository
}

func NewTimetableService(r *repository.TimetableRepository) *TimetableService {
	return &TimetableService{repo: r}
}

func (s *TimetableService) GetTimetable(userID string, year int, semester string) ([]model.TimetableItem, error) {
	if year == 0 {
		year = time.Now().Year()
	}

	if semester == "" {
		semester = "1st"
	}

	return s.repo.GetTimetableByUserID(userID, year, semester)
}
