package services

import (
	"context"
	"errors"
	"reka-storage/src/user/models"
	"reka-storage/src/user/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Register(
	ctx context.Context,
	username, email, password string,
) (*models.User, error) {

	if _, err := s.repo.FindByEmail(ctx, email); err == nil {
		return nil, errors.New("email already registered")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := s.repo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.Password = ""
	return createdUser, nil

}
