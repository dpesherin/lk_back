package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
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
		fmt.Println("Ошибка при преобразовании:", err)
		return nil, err
	}
	userModel := &models.User{
		ID: int64(uid),
	}
	fmt.Println(userModel.ID)
	user, err := us.ur.GetUserById(userModel.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) CreateUser(ctx *gin.Context) (*models.User, error) {
	u := &models.UserData{}
	err := ctx.ShouldBindJSON(u)
	fmt.Println(u)
	if err != nil {
		return nil, err
	}
	user, err := us.ur.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}
