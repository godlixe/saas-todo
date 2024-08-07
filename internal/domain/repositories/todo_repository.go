package repositories

import (
	"context"
	"saas-todo-list/internal/domain/entities"

	"github.com/google/uuid"
)

type TodoRepository interface {
	GetAll(
		ctx context.Context,
		filter entities.TodoFilter,
	) ([]entities.Todos, error)

	FindById(
		ctx context.Context,
		id uuid.UUID,
		tenantId uuid.UUID,
	) (*entities.Todos, error)

	CreateTodo(
		ctx context.Context,
		todo *entities.Todos,
	) error

	UpdateTodo(
		ctx context.Context,
		todo *entities.Todos,
	) error

	DeleteTodo(
		ctx context.Context,
		id uuid.UUID,
		tenantId uuid.UUID,
	) error
}
