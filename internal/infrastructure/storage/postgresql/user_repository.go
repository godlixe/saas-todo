package postgresql

import (
	"context"
	"saas-todo-list/internal/domain/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	ctx context.Context,
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindById(
	ctx context.Context,
	id uuid.UUID,

	// tenant identifier
	tenantId uuid.UUID,
) (*entities.User, error) {
	var user entities.User

	tx := r.db.Model(entities.User{}).Where("id = ? AND tenant_id = ?", id, tenantId).First(&user)

	return &user, tx.Error
}

func (r *UserRepository) FindByUsername(
	ctx context.Context,
	username string,

	// tenant identifier
	tenantId uuid.UUID,
) (*entities.User, error) {
	var user entities.User

	tx := r.db.Model(entities.User{}).Where("username = ? AND tenant_id = ?", username, tenantId).First(&user)

	return &user, tx.Error
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	user *entities.User,
) error {
	tx := r.db.Model(entities.User{}).Create(&user)

	return tx.Error
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	user *entities.User,
) error {
	tx := r.db.Model(entities.User{}).Updates(&user)

	return tx.Error
}
