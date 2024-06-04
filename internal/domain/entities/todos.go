package entities

import "github.com/google/uuid"

type Todos struct {
	ID uuid.UUID `json:"id"`

	// Owner of the Todo
	UserID uuid.UUID `json:"user_id"`
	User   *User     `json:"user"`

	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`

	TimeStamp
	TenantData
}
