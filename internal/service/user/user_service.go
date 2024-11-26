package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/models/special_models"
	users_repo "lk_back/internal/repository/users"
	"strconv"
)

type UserService struct {
	ur users_repo.UserRepoInterface
}

func NewUserService(ur *users_repo.UserRepo) *UserService {
	return &UserService{
		ur: ur,
	}
}

func (us *UserService) GetUserById(ctx *gin.Context) (*models.User, error) {
	uid, err := strconv.Atoi(ctx.Param("id")) // Преобразуем строку в целое число
	if err != nil {
		return nil, errors.New("id isn't number")
	}
	userModel := &models.User{
		ID: int64(uid),
	}
	user, err := us.ur.GetUserById(userModel.ID)
	if err != nil {
		return nil, errors.New("user wasn't found")
	}
	return user, nil
}

func (us *UserService) CreateUser(ctx *gin.Context) (*models.User, error) {
	u := &special_models.UserData{}
	err := ctx.ShouldBindJSON(u)
	if err != nil {
		return nil, errors.New("invalid request format")
	}
	user, err := us.ur.CreateUser(u)
	if err != nil {
		return nil, errors.New("error while creating user")
	}
	return user, nil
}
