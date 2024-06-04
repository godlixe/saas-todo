package controllers

import (
	"net/http"
	"saas-todo-list/internal/domain/entities"
	"saas-todo-list/internal/presentation/dto"
	"saas-todo-list/internal/presentation/services"
	"saas-todo-list/pkg/httpx"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(
	authService services.AuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var params dto.LoginDto

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.authService.Login(ctx, &entities.User{
		Username: params.Username,
		Password: params.Password,
		TenantData: entities.TenantData{
			TenantID: uuid.MustParse(params.TenantId),
		},
	})
	if err != nil {
		err = httpx.NewError("error logging in", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "login successfull",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: res,
		},
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var params dto.RegisterDto

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.authService.Register(ctx, &entities.User{
		Name:     params.Name,
		Username: params.Username,
		Password: params.Password,
		TenantData: entities.TenantData{
			TenantID: uuid.MustParse(params.TenantId),
		},
	})
	if err != nil {
		err = httpx.NewError("error registering", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "register successfull",
		Data:    res,
	})
}
