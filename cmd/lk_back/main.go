package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"lk_back/internal/config"
	"lk_back/internal/pkg/server/router"
	users_repo "lk_back/internal/repository/users"
	"lk_back/internal/service/auth"
	"lk_back/internal/service/user"
	"log"
)

func main() {
	conf, err := config.LoadConfig("../../config/config.yaml")
	if err != nil {
		log.Fatalf("Error while loading conf file: \n %v", err)
	}
	conConf := pgx.ConnConfig{
		User:     conf.Database.User,
		Database: conf.Database.DB,
		Password: conf.Database.Password,
		Host:     conf.Database.Host,
		Port:     uint16(conf.Database.Port),
	}
	poolConf := pgx.ConnPoolConfig{
		ConnConfig:     conConf,
		MaxConnections: 2,
	}
	db, err := pgx.NewConnPool(poolConf)
	if err != nil {
		log.Fatalf("Error while create PoolConn: \n %v", err)
	}

	r := gin.Default()

	//USERS
	ur := users_repo.NewUserRepo(db)
	us := user.NewUserService(ur)
	uRouter := router.NewUserRouter(r, us)
	uRouter.SetupRouter()
	//END USERS
	//AUTH
	as := auth.NewAuthService(ur)
	aRouter := router.NewAuthRouter(r, as)
	aRouter.SetupRoutes()
	//END AUTH
	r.Run(":" + conf.Server.Port)
}
