package service

import (
	"errors"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type ProfileService struct {
	userRepo    *repository.UserRepository
	studentRepo *repository.StudentRepository
}

func NewProfileService(userRepo *repository.UserRepository, studentRepo *repository.StudentRepository) *ProfileService {
	return &ProfileService{
		userRepo:    userRepo,
		studentRepo: studentRepo,
	}
}

func (s *ProfileService) GetProfile(userID string) (*model.Profile, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("ユーザーが見つかりません")
	}

	profile := &model.Profile{
		Name:          user.Name,
		StudentID:     "",
		Email:         user.Email,
		Phone:         "",
		Department:    "",
		Grade:         "",
		EnrollmentYear: "",
		Birthday:      "",
	}

	return profile, nil
}

func (s *ProfileService) GetProfileByStudentNumber(studentNumber string) (*model.Profile, error) {
	student, err := s.studentRepo.GetByStudentNumber(studentNumber)
	if err != nil {
		return nil, errors.New("学生情報が見つかりません")
	}

	profile := &model.Profile{
		Name:          student.Name,
		StudentID:     student.StudentNumber,
		Email:         "",
		Phone:         "",
		Department:    student.Department,
		Grade:         student.Grade,
		EnrollmentYear: "",
		Birthday:      "",
	}

	return profile, nil
}

