package services

import (
	"context"
	"saas-todo-list/internal/domain/entities"
)

type AuthService interface {
	Login(
		ctx context.Context,
		userRequest *entities.User,
	) (string, error)

	Register(
		ctx context.Context,
		user *entities.User,
	) (*entities.User, error)
}
