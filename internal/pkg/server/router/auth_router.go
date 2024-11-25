package router

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/service/auth"
	"net/http"
)

type AuthRouter struct {
	r  *gin.Engine
	as *auth.AuthService
}

func NewAuthRouter(r *gin.Engine, as *auth.AuthService) *AuthRouter {
	return &AuthRouter{
		r:  r,
		as: as,
	}
}

func (ar *AuthRouter) SetupRoutes() {
	group := ar.r.Group("/auth")
	group.POST("/login", func(ctx *gin.Context) {
		jwt, err := ar.as.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: jwt, Message: ""})
		return
	})
}
