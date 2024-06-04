package entities

import (
	"saas-todo-list/pkg/auth"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"-"`

	TimeStamp
	TenantData
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	hashedPassword, err := auth.HashAndSalt(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}
