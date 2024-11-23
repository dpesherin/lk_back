package httplayer

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/applayer/services"
	"log"
	"net/http"
	"strconv"
)

type UserRouter struct {
	engine *gin.Engine
	route  *gin.RouterGroup
}

func (ur *UserRouter) Init() *UserRouter {
	router := ur.engine.Group("/user")

	router.GET("/:id", func(ctx *gin.Context) {
		us := &services.UserService{Ctx: ctx}
		uid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Fatalf("ID пользователя не является числом: \n %v", err)
			return
		}
		user, err := us.GetUser(int64(uid))
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, &user)
	})
	ur.route = router
	return ur
}
