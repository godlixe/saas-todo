package middlewares

import (
	"net/http"
	"saas-todo-list/pkg/httpx"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "unauthorized user",
			})
			return
		}

		tokenHeader := strings.Split(token, "Bearer ")

		if len(tokenHeader) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "corrupt header format",
			})
			return
		}

		// claims := auth.Claims{}
		// jwtToken, err := jwt.ParseWithClaims(tokenHeader[1], &claims, func(token *jwt.Token) (interface{}, error) {
		// 	return auth.JWTKey, nil
		// })
		// fmt.Println(auth.JWTKey)
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
		// 		Message: "error parsing token",
		// 	})
		// 	return
		// }

		// if !jwtToken.Valid {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
		// 		Message: "invalid token",
		// 	})
		// 	return
		// }

		tenantID := ctx.GetHeader("x-tenant-id")
		if tenantID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "invalid tenant_id",
			})
			return
		}

		userID := ctx.GetHeader("x-user-id")
		if userID == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "invalid user_id",
			})
			return
		}

		ctx.Set("user_id", userID)
		ctx.Set("tenant_id", tenantID)
		ctx.Next()
	}
}
