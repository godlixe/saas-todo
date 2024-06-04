package services

import (
	"context"
	"saas-todo-list/internal/domain/entities"
	"saas-todo-list/internal/domain/repositories"

	"github.com/google/uuid"
)

type TodoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(
	todoRepository repositories.TodoRepository,
) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (s *TodoService) FindById(
	ctx context.Context,
	id uuid.UUID,
	tenantId uuid.UUID,
) (*entities.Todos, error) {
	return s.todoRepository.FindById(ctx, id, tenantId)
}

func (s *TodoService) CreateTodo(
	ctx context.Context,
	todo *entities.Todos,
) (*entities.Todos, error) {
	return todo, s.todoRepository.CreateTodo(ctx, todo)
}

func (s *TodoService) UpdateTodo(
	ctx context.Context,
	todo *entities.Todos,
) (*entities.Todos, error) {
	return todo, s.todoRepository.UpdateTodo(ctx, todo)
}

func (s *TodoService) DeleteTodo(
	ctx context.Context,
	id uuid.UUID,
	tenantId uuid.UUID,
) error {
	return s.todoRepository.DeleteTodo(ctx, id, tenantId)
}
