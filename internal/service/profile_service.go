package service

import (
	"database/sql"
	"errors"

	"github.com/oita-u/campus-api/internal/logger"
	"github.com/oita-u/campus-api/internal/model"
	"github.com/oita-u/campus-api/internal/repository"
	"go.uber.org/zap"
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
		logger.Get().Error("Error retrieving profile", zap.String("userID", userID), zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Profile{}, nil
		}

		return nil, errors.New("ユーザーが見つかりません")
	}
	return profile, nil
}
