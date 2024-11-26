package middleware

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/service/jwt"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.Request.Header.Get("Authorization")
		user, err := jwt.ValidateToken(jwtToken)
		if err != nil {
			ctx.JSON(http.StatusForbidden, &models.Response{
				Success: false,
				Message: err.Error(),
				Obj:     nil,
			})
			ctx.Abort()
			return
		}
		ctx.Set("decodedToken", user)
		ctx.Next()
	}
}
