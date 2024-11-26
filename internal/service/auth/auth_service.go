package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lk_back/internal/models/special_models"
	users_repo "lk_back/internal/repository/users"
	jwt2 "lk_back/internal/service/jwt"
)

type AuthService struct {
	ur users_repo.UserRepoInterface
}

type ClaimsAuth struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

func NewAuthService(ur users_repo.UserRepoInterface) *AuthService {
	return &AuthService{
		ur: ur,
	}
}

func (as *AuthService) Login(ctx *gin.Context) (*jwt2.JWT, error) {
	claims := &ClaimsAuth{}
	err := ctx.ShouldBindJSON(claims)
	if err != nil {
		return nil, errors.New("error parsing claims")
	}
	candData, err := as.ur.GetUserByLogin(claims.Login)
	if err != nil {
		return nil, errors.New("user wasn't found")
	}
	fmt.Println(claims.Pass)
	err = bcrypt.CompareHashAndPassword([]byte(candData.Password), []byte(claims.Pass))
	if err != nil {
		return nil, errors.New("user wasn't found")
	}
	cand := &special_models.TokenData{
		ID:    candData.ID,
		Login: candData.Login,
		Admin: candData.Admin,
	}
	jwt, err := jwt2.GeneratePair(cand)
	if err != nil {
		return nil, errors.New("error while creating JWT")
	}
	return jwt, nil
}

type PasswordChangeRequest struct {
	Login    string `json:"login"`
	Password string `json:"password" binding:"required"`
}

func (as *AuthService) ChangePassword(ctx *gin.Context) error {
	u := ctx.MustGet("decodedToken").(*special_models.TokenData)
	rq := &PasswordChangeRequest{}
	var login string
	err := ctx.ShouldBindJSON(rq)
	if err != nil {
		return errors.New("invalid request format")
	}
	if u.Admin {
		login = rq.Login
	} else {
		if u.Login == rq.Login {
			login = u.Login
		} else {
			return errors.New("user isn't admin")
		}
	}
	err = as.ur.ChangePassword(login, rq.Password)
	if err != nil {
		return errors.New("error while changing password")
	}
	return nil
}
