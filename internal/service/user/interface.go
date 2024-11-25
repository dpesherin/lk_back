package user

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
)

type UserServiceInterface interface {
	GetUserById(ctx *gin.Context) (*models.User, error)
	CreateUser(ctx *gin.Context) (*models.User, error)
}
