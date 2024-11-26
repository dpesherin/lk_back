package router

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/pkg/server/middleware"
	"lk_back/internal/service/auth"
	"net/http"
)

type AuthRouter struct {
	r  *gin.Engine
	as auth.AuthInterface
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
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: "Auth Error"})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: jwt, Message: ""})
		return
	})
	group.POST("/changepwd", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		err := ar.as.ChangePassword(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: "Error while changing password"})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: nil, Message: ""})
		return
	})
}
