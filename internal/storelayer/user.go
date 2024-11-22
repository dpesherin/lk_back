package storelayer

import (
	"fmt"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
	"lk_back/internal/models"
	"log"
)

func (store *store) CreateUser(login string, email string, pass string) error {
	user := models.User{Login: login, Active: true, Email: email}
	conn, err := pgx.Connect(store.dbConfig)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД\n %v", err)
	}
	defer func(conn *pgx.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Ошибка при разрыве соединения с БД\n %v", err)
		}
	}(conn)
	query := `
        INSERT INTO users (Login, Active, Email, HashPass)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	var id int64
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка при создании пользователя:\n %v", err)
		return err
	}
	err = conn.QueryRow(query, user.Login, user.Active, user.Email, hash).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Printf("Новый пользователь добавлен с ID: %d\n", id)
	return nil
}

func (store *store) GetUser(id int64) (*models.User, error) {
	user := models.User{}
	conn, err := pgx.Connect(store.dbConfig)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД\n %v", err)
	}
	defer func(conn *pgx.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Ошибка при разрыве соединения с БД\n %v", err)
		}
	}(conn)
	query := `
        SELECT * FROM users
        WHERE ID=$1
    `
	err = conn.QueryRow(query, id).Scan(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}
