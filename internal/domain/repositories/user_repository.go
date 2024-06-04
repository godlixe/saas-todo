package repositories

import (
	"context"
	"saas-todo-list/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindById(
		ctx context.Context,
		id uuid.UUID,

		// tenant identifier
		tenantId uuid.UUID,
	) (*entities.User, error)

	FindByUsername(
		ctx context.Context,
		username string,

		// tenant identifier
		tenantId uuid.UUID,
	) (*entities.User, error)

	CreateUser(
		ctx context.Context,
		user *entities.User,
	) error

	UpdateUser(
		ctx context.Context,
		user *entities.User,
	) error
}
