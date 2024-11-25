package users_repo

import "lk_back/internal/models"

type UserRepoInterface interface {
	GetUserById(id int64) (*models.User, error)
	GetUserByLogin(login string) (*models.UserData, error)
	CreateUser(u *models.UserData) (*models.User, error)
}
