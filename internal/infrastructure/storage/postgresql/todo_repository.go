package postgresql

import (
	"context"
	"saas-todo-list/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(
	ctx context.Context,
	db *gorm.DB,
) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) FindById(
	ctx context.Context,
	id uuid.UUID,

	// tenant identifier
	tenantId uuid.UUID,
) (*entities.Todos, error) {
	var todo entities.Todos

	tx := r.db.Model(entities.Todos{}).Where("id = ? AND tenant_id = ?", id, tenantId).First(&todo)

	return &todo, tx.Error
}

func (r *TodoRepository) CreateTodo(
	ctx context.Context,
	todo *entities.Todos,
) error {
	tx := r.db.Model(entities.Todos{}).Create(&todo)

	return tx.Error
}

func (r *TodoRepository) UpdateTodo(
	ctx context.Context,
	todo *entities.Todos,
) error {
	tx := r.db.Model(entities.Todos{}).Updates(&todo)

	return tx.Error
}

func (r *TodoRepository) DeleteTodo(
	ctx context.Context,
	id uuid.UUID,
	tenantId uuid.UUID,
) error {
	tx := r.db.Model(entities.Todos{}).Where("id = ? AND tenant_id = ?", id, tenantId).Delete(entities.Todos{})

	return tx.Error
}
