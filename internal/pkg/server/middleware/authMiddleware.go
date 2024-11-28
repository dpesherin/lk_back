package middleware

import (
	"github.com/gin-gonic/gin"
	"lk_back/internal/models"
	"lk_back/internal/models/special_models"
	"lk_back/internal/service/jwt"
	"net/http"
)

// TODO: Complete AuthForeign + RefreshAuth
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken, err := ctx.Cookie("accessToken")
		if err != nil {
			ctx.JSON(http.StatusForbidden, &models.Response{
				Success: false,
				Message: "unauthorized user",
				Obj:     nil,
			})
			ctx.Abort()
			return
		}
		refreshToken, err := ctx.Cookie("refreshToken")
		if err != nil {
			ctx.JSON(http.StatusForbidden, &models.Response{
				Success: false,
				Message: "unauthorized user",
				Obj:     nil,
			})
			ctx.Abort()
			return
		}
		foreignToken, err := ctx.Cookie("foreignToken")
		if err != nil {
			user, err := tryToReconnect(ctx, authToken, refreshToken)
			if err != nil {
				ctx.JSON(http.StatusForbidden, &models.Response{
					Success: false,
					Message: "unauthorized user",
					Obj:     nil,
				})
				ctx.Abort()
				return
			}
			ctx.Set("decodedToken", user)
			ctx.Next()
		} else {
			user, err := jwt.ValidateToken(foreignToken)
			if err != nil {
				user, err := tryToReconnect(ctx, authToken, refreshToken)
				if err != nil {
					ctx.JSON(http.StatusForbidden, &models.Response{
						Success: false,
						Message: "unauthorized user",
						Obj:     nil,
					})
					ctx.Abort()
					return
				}
				ctx.Set("decodedToken", user)
				ctx.Next()
			}
			ctx.Set("decodedToken", user)
			ctx.Next()
		}

	}
}

func tryToReconnect(ctx *gin.Context, accessToken string, refreshToken string) (*special_models.TokenData, error) {
	user, err := jwt.ValidateToken(accessToken)
	if err != nil {
		user, err := jwt.ValidateToken(refreshToken)
		if err != nil {
			return nil, err
		}
		jwtPair, err := jwt.GeneratePair(user)
		if err != nil {
			return nil, err
		}
		ctx.SetCookie("accessToken", jwtPair.AccessToken, 10800, "/", "192.168.0.100", false, false)
		ctx.SetCookie("refreshToken", jwtPair.RefreshToken, 86400, "/", "192.168.0.100", false, false)
		return user, nil
	}
	return user, nil
}
