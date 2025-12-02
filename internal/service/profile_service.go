package service

import (
	"errors"

	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
)

type ProfileService struct {
	profileRepo *repository.ProfileRepository
}

func NewProfileService(profileRepo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepo: profileRepo,
	}
}

func (s *ProfileService) GetProfile(userID string) (*model.Profile, error) {
	profile, err := s.profileRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("ユーザーが見つかりません")
	}
	return profile, nil
}
