package service

import (
	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(r *repository.CourseRepository) *CourseService {
	return &CourseService{repo: r}
}

func (s *CourseService) SearchCourses(userID string, year int, semester, keyword string) ([]model.Course, int, string, error) {
	year, semester = resolveYearSemester(year, semester)

	courses, err := s.repo.SearchCourses(userID, year, semester, keyword)
	if err != nil {
		return nil, 0, "", err
	}

	return courses, year, semester, nil
}
