package applayer

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
)

type UserService struct {
	Ctx *gin.Context
}

func (us *UserService) GetUser(id int64) (*models.User, error) {
	return &models.User{ID: int64(id), Login: "admin"}, nil
}
