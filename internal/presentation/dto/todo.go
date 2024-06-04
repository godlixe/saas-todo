package dto

type CreateTodoDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTodoDto struct {
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}
