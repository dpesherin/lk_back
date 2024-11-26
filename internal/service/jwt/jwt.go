package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"lk_back/internal/models/special_models"
	"strings"
	"time"
)

type JWT struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type Claims struct {
	User special_models.TokenData `json:"user"`
	jwt.StandardClaims
}

var secretKeyPhrase = "SuperS3cr3tK3y"

func GenerateToken(u *special_models.TokenData) (string, error) {
	claims := &Claims{
		User: *u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // Токен действителен в течение 24 часов
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(secretKeyPhrase)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*special_models.TokenData, error) {
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid token format")
	}

	tokenStr := parts[1]

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKeyPhrase), nil
	})

	if err != nil {
		var validationErr *jwt.ValidationError
		if errors.As(err, &validationErr) {
			switch validationErr.Errors {
			case jwt.ValidationErrorMalformed:
				return nil, errors.New("invalid token format")
			case jwt.ValidationErrorExpired:
				return nil, errors.New("token has expired")
			default:
				return nil, errors.New("couldn't handle this token")
			}
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &claims.User, nil
}

func GeneratePair(u *special_models.TokenData) (*JWT, error) {
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
