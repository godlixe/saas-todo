package services

import (
	"context"
	"saas-todo-list/internal/domain/entities"
	"saas-todo-list/internal/domain/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) FindById(
	ctx context.Context,
	id uuid.UUID,
	tenantId uuid.UUID,
) (*entities.User, error) {
	return s.userRepository.FindById(ctx, id, tenantId)
}

func (s *UserService) CreateUser(
	ctx context.Context,
	user *entities.User,
) (*entities.User, error) {
	return user, s.userRepository.CreateUser(ctx, user)
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	user *entities.User,
) (*entities.User, error) {
	return user, s.userRepository.UpdateUser(ctx, user)
}
