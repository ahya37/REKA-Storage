package services

import (
	"context"
	"reka-storage/src/user/models"
	"reka-storage/src/user/repositories"
)

type Service struct {
	repo *repositories.UserRepository
}

func NewService(repo *repositories.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProfile(ctx context.Context, userID string) (*models.User, error) {
	return s.repo.FindByID(ctx, userID)
}
