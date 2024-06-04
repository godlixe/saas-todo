package services

import (
	"context"
	"saas-todo-list/internal/domain/entities"

	"github.com/google/uuid"
)

type UserService interface {
	FindById(
		ctx context.Context,
		id uuid.UUID,
	) (*entities.User, error)

	CreateUser(
		ctx context.Context,
		user *entities.User,
	) (*entities.User, error)

	UpdateUser(
		ctx context.Context,
		user *entities.User,
	) (*entities.User, error)
}
