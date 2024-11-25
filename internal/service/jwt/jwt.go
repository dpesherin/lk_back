package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"lk_back/internal/models"
	"time"
)

type JWT struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type Claims struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

func GenerateToken(u *models.User) (string, error) {
	claims := &Claims{
		User: *u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // Токен действителен в течение 24 часов
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("SuperS3cr3tK3y") // Замените на свой секретный ключ

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte("SuperS3cr3tK3y")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GeneratePair(u *models.User) (*JWT, error) {
	access, err := GenerateToken(u)
	if err != nil {
		return nil, err
	}
	refresh, err := GenerateToken(u)
	if err != nil {
		return nil, err
	}
	jwtpair := &JWT{}
	jwtpair.AccessToken = access
	jwtpair.RefreshToken = refresh
	return jwtpair, nil
}
