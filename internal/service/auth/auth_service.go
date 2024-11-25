package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"lk_back/internal/models"
	users_repo "lk_back/internal/repository/users"
	jwt2 "lk_back/internal/service/jwt"
)

type AuthService struct {
	ur *users_repo.UserRepo
}

type ClaimsAuth struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func NewAuthService(ur *users_repo.UserRepo) *AuthService {
	return &AuthService{
		ur: ur,
	}
}

func (as *AuthService) Login(ctx *gin.Context) (*jwt2.JWT, error) {
	claims := &ClaimsAuth{}
	err := ctx.ShouldBindJSON(claims)
	if err != nil {
		return nil, err
	}
	fmt.Println(claims.Login)
	candData, err := as.ur.GetUserByLogin(claims.Login)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(candData.Password), []byte(claims.Pass))
	if err != nil {
		return nil, err
	}
	cand := &models.User{
		ID:       candData.ID,
		Login:    candData.Login,
		Active:   candData.Active,
		Email:    candData.Email,
		Name:     candData.Name,
		LastName: candData.LastName,
		Avatar:   candData.Avatar,
		Admin:    candData.Admin,
	}
	jwt, err := jwt2.GeneratePair(cand)
	if err != nil {
		return nil, err
	}
	return jwt, nil
}
