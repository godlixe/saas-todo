package services

import (
	"context"
	"errors"
	"saas-todo-list/internal/domain/entities"
	"saas-todo-list/internal/domain/repositories"
	"saas-todo-list/pkg/auth"

	"gorm.io/gorm"
)

type AuthService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(
	userRepository repositories.UserRepository,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

// Login returns an access token if successfull
// and returns an error otherwise.
func (s *AuthService) Login(
	ctx context.Context,
	userRequest *entities.User,
) (string, error) {
	user, err := s.userRepository.FindByUsername(ctx, userRequest.Username, userRequest.TenantID)
	if err != nil {
		return "", err
	}

	isPasswordMatch, err := auth.ComparePassword(
		user.Password,
		[]byte(userRequest.Password),
	)
	if err != nil {
		return "", err
	}

	if !isPasswordMatch {
		return "", errors.New("password does not match")
	}

	token, err := auth.GenerateJWTToken(
		user.ID.String(),
		user.TenantID.String(),
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login returns an access token if successfull
// and returns an error otherwise.
func (s *AuthService) Register(
	ctx context.Context,
	userRequest *entities.User,
) (*entities.User, error) {
	_, err := s.userRepository.FindByUsername(ctx, userRequest.Username, userRequest.TenantID)
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return nil, err
	}

	if err == nil {
		return nil, errors.New("user with same username exists")
	}

	err = s.userRepository.CreateUser(ctx, userRequest)
	if err != nil {
		return nil, err
	}

	return userRequest, nil
}
