package controllers

import (
	"errors"
	"net/http"
	"saas-todo-list/internal/domain/entities"
	"saas-todo-list/internal/presentation/dto"
	"saas-todo-list/internal/presentation/services"
	"saas-todo-list/pkg/httpx"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(
	todoService services.TodoService,
) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (c *TodoController) GetAll(ctx *gin.Context) {
	getPersonal := ctx.Query("personal")

	tenantId, ok := ctx.Get("tenant_id")
	if !ok {
		err := httpx.NewError("tenant unidentifiable", errors.New("tenant unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	strTenantId := tenantId.(string)

	filter := entities.TodoFilter{
		TenantID: strTenantId,
	}

	if getPersonal == "1" {
		filter.UserID = ctx.GetString("user_id")
	}

	res, err := c.todoService.GetAll(
		ctx,
		filter,
	)
	if err != nil {
		err := httpx.NewError("todos not found", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get todos successfull",
		Data:    res,
	})
}

func (c *TodoController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	tenantId, ok := ctx.Get("tenant_id")
	if !ok {
		err := httpx.NewError("tenant unidentifiable", errors.New("tenant unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	strTenantId := tenantId.(string)

	if err := uuid.Validate(id); err != nil {
		err := httpx.NewError("todo id unidentifiable", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	res, err := c.todoService.FindById(
		ctx,
		uuid.MustParse(id),
		uuid.MustParse(strTenantId),
	)
	if err != nil {
		err := httpx.NewError("tenant id not found", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get todo successfull",
		Data:    res,
	})
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var (
		todo dto.CreateTodoDto
		err  error
	)

	tenantId, ok := ctx.Get("tenant_id")
	if !ok {
		err := httpx.NewError("tenant unidentifiable", errors.New("tenant unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	userId, ok := ctx.Get("user_id")
	if !ok {
		err := httpx.NewError("user unidentifiable", errors.New("user unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	strTenantId := tenantId.(string)
	strUserId := userId.(string)

	err = ctx.Bind(&todo)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.todoService.CreateTodo(
		ctx,
		&entities.Todos{
			ID:     uuid.New(),
			Name:   todo.Name,
			UserID: uuid.MustParse(strUserId),
			TenantData: entities.TenantData{
				TenantID: uuid.MustParse(strTenantId),
			},
			Content: todo.Content,
		},
	)
	if err != nil {
		err := httpx.NewError("failed to create todo", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "create todo successfull",
		Data:    res,
	})
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	var (
		todo dto.UpdateTodoDto
		err  error
	)

	id := ctx.Param("id")

	tenantId, ok := ctx.Get("tenant_id")
	if !ok {
		err := httpx.NewError("tenant unidentifiable", errors.New("tenant unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	strTenantId := tenantId.(string)

	err = ctx.Bind(&todo)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.todoService.UpdateTodo(
		ctx,
		&entities.Todos{
			ID:     uuid.MustParse(id),
			Name:   todo.Name,
			IsDone: &todo.IsDone,
			TenantData: entities.TenantData{
				TenantID: uuid.MustParse(strTenantId),
			},
			Content: todo.Content,
		},
	)
	if err != nil {
		err := httpx.NewError("failed to update todo", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "update todo successfull",
		Data:    res,
	})
}

func (c *TodoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	tenantId, ok := ctx.Get("tenant_id")
	if !ok {
		err := httpx.NewError("tenant unidentifiable", errors.New("tenant unidentifiable"), http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	strTenantId := tenantId.(string)

	err := c.todoService.DeleteTodo(
		ctx,
		uuid.MustParse(id),
		uuid.MustParse(strTenantId),
	)
	if err != nil {
		err := httpx.NewError("tenant id not found", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "delete todo successfull",
		Data:    true,
	})
}
