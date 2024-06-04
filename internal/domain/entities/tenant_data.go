package entities

import uuid "github.com/google/uuid"

type TenantData struct {
	TenantID uuid.UUID `json:"tenant_id"`
}
