package entities

import "github.com/google/uuid"

type Todos struct {
	ID uuid.UUID `json:"id"`

	// Owner of the Todo
	UserID uuid.UUID `json:"user_id"`
	User   *User     `json:"user,omitempty"`

	Name    string `json:"name"`
	IsDone  *bool  `json:"is_done"`
	Content string `json:"content"`

	TimeStamp
	TenantData
}

type TodoFilter struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
}
