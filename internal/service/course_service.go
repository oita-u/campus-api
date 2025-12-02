package service

import (
	"time"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(r *repository.CourseRepository) *CourseService {
	return &CourseService{repo: r}
}

func (s *CourseService) SearchCourses(year int, semester, keyword string) ([]model.Course, int, string, error) {
	if year == 0 {
		year = time.Now().Year()
	}

	if semester == "" {
		semester = "1st"
	}

	courses, err := s.repo.SearchCourses(year, semester, keyword)
	if err != nil {
		return nil, 0, "", err
	}

	return courses, year, semester, nil
}
