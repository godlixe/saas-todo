package dto

type CreateTodoDto struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateTodoDto struct {
	Name    string `json:"name"`
	IsDone  bool   `json:"is_done"`
	Content string `json:"content"`
}
