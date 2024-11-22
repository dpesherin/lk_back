package httplayer

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/config"
	"log"
	"net/http"
)

type HttpServer struct {
	engine *gin.Engine
}

func New(conf *config.Config) (*HttpServer, error) {
	r := gin.Default()
	server := HttpServer{r}
	server.SetupRoutes()
	err := r.Run(":" + conf.Server.Port)
	if err != nil {
		log.Fatalf("Ошибка инициализации сервера: \n %v", err)
		return &server, nil
	}
	return &server, nil
}

func (server *HttpServer) SetupRoutes() {
	server.engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "INIT")
	})
}
