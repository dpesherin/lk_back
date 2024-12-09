package router

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/interfaces"
	"lk_back/internal/models"
	"lk_back/internal/pkg/server/middleware"
	"net/http"
)

type AuthRouter struct {
	r  *gin.Engine
	as interfaces.AuthServiceInterface
}

func NewAuthRouter(r *gin.Engine, as interfaces.AuthServiceInterface) *AuthRouter {
	return &AuthRouter{
		r:  r,
		as: as,
	}
}

func (ar *AuthRouter) SetupRoutes() {
	group := ar.r.Group("/auth")
	group.POST("/login", func(ctx *gin.Context) {
		jwtPair, err := ar.as.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: err.Error()})
			return
		}
		ctx.SetCookie("accessToken", jwtPair.AccessToken, 10800, "/", "", true, true)
		ctx.SetCookie("refreshToken", jwtPair.RefreshToken, 86400, "/", "", true, true)
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: jwtPair, Message: ""})
		return
	})
	group.POST("/changepwd", middleware.AuthMiddleware(), func(ctx *gin.Context) {
		err := ar.as.ChangePassword(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: nil, Message: ""})
		return
	})
}
