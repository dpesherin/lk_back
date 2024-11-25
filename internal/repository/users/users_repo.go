package users_repo

import (
	"fmt"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"lk_back/internal/models"
)

type UserRepo struct {
	db *pgx.ConnPool
}

func NewUserRepo(db *pgx.ConnPool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) GetUserById(id int64) (*models.User, error) {
	user := &models.User{}
	err := ur.db.QueryRow(`
		SELECT ID, LOGIN, ACTIVE, EMAIL, NAME, LAST_NAME, AVATAR, ADMIN FROM users WHERE ID=$1
	`, id).Scan(&user.ID, &user.Login, &user.Active, &user.Email, &user.Name, &user.LastName, &user.Avatar, &user.Admin)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) GetUserByLogin(login string) (*models.UserData, error) {
	user := &models.UserData{}
	err := ur.db.QueryRow(`
		SELECT * FROM users WHERE LOGIN=$1;
	`, login).Scan(&user.ID, &user.Login, &user.Active, &user.Email, &user.Name, &user.LastName, &user.Avatar, &user.Admin, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) CreateUser(u *models.UserData) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	fmt.Println(u)
	err = ur.db.QueryRow(`
		INSERT INTO users (LOGIN, ACTIVE, EMAIL, NAME, LAST_NAME, AVATAR, ADMIN, PASS)
    	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING ID, LOGIN, ACTIVE, EMAIL, NAME, LAST_NAME, AVATAR, ADMIN;
	`, u.Login, u.Active, u.Email, u.Name, u.LastName, u.Avatar, u.Admin, string(hash)).Scan(&user.ID, &user.Login, &user.Active, &user.Email, &user.Name, &user.LastName, &user.Avatar, &user.Admin)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return user, nil
}
