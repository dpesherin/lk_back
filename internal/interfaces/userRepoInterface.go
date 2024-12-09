package interfaces

import (
	"lk_back/internal/models"
	"lk_back/internal/models/special_models"
)

type UserRepoInterface interface {
	GetUserById(id int64) (*models.User, error)
	GetUserByLogin(login string) (*special_models.UserData, error)
	CreateUser(u *special_models.UserData) (*models.User, error)
	ChangePassword(login string, password string) error
}
