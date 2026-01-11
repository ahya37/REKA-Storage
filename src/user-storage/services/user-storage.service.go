package services

import (
	"context"
	"reka-storage/src/user-storage/models"
	"reka-storage/src/user-storage/repositories"
)

type UserStorageService struct {
	repo *repositories.UserStorageRepository
}

func NewUserStorageService(repo *repositories.UserStorageRepository) *UserStorageService {
	return &UserStorageService{repo: repo}
}

func (s *UserStorageService) GetUsage(ctx context.Context, userID string) (*models.UserStorage, error) {
	return s.repo.GetUsage(
		ctx,
		userID,
	)
}
