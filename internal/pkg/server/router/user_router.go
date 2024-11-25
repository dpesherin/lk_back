package router

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/service/user"
	"net/http"
)

type UserRouter struct {
	r  *gin.Engine
	us *user.UserService
}

func NewUserRouter(r *gin.Engine, us *user.UserService) *UserRouter {
	return &UserRouter{
		r:  r,
		us: us,
	}
}

func (ur *UserRouter) SetupRouter() {
	group := ur.r.Group("/user")
	group.GET("/:id", func(ctx *gin.Context) {
		u, err := ur.us.GetUserById(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: u, Message: ""})
		return
	})

	group.POST("/create", func(ctx *gin.Context) {
		u, err := ur.us.CreateUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &models.Response{Success: false, Obj: nil, Message: "Error while creating User"})
			return
		}
		ctx.JSON(http.StatusOK, &models.Response{Success: true, Obj: u, Message: "Success creating"})
		return
	})
}
