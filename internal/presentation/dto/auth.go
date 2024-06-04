package dto

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	TenantId string `json:"tenant_id" binding:"required"`
}

type RegisterDto struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	TenantId string `json:"tenant_id" binding:"required"`
}
