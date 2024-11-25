package auth

import (
	"github.com/gin-gonic/gin"
	jwt2 "lk_back/internal/service/jwt"
)

type AuthInterface interface {
	Login(ctx *gin.Context) (*jwt2.JWT, error)
}
