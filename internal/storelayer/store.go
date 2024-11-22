package storelayer

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"lk_back/internal/models"
)

type store struct {
	dbConfig pgx.ConnConfig
}

type Store interface {
	CreateUser(ctx gin.Context) error
	GetUser(ctx gin.Context) (models.User, error)
}
