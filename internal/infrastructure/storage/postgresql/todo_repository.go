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

func (r *TodoRepository) GetAll(
	ctx context.Context,
	filter entities.TodoFilter,
) ([]entities.Todos, error) {
	var todos []entities.Todos

	tx := r.db.Model(&entities.Todos{})

	if filter.UserID != "" {
		tx.Where("user_id = ?", filter.UserID)
	}

	tx = tx.Where("tenant_id = ?", filter.TenantID).Find(&todos)

	return todos, tx.Error
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
	tx := r.db.Model(entities.Todos{}).Where("id = ?", todo.ID).Updates(&todo)

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
