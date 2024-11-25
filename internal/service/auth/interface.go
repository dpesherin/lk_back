package auth

import "lk_back/internal/models"

type AuthServiceInterface interface {
	Authenticate(username string, password string) (user *models.User, jwt string, err error)
	ChangePassword(userID int64, oldPassword, newPassword string) error
}
