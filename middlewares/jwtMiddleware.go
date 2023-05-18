package middlewares

import (
	"backend-golang/constants"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateToken(email string, role string, id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["user_id"] = id
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_KEY))
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["role"].(string)
		if isAdmin != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isUser := claims["role"].(string)
		if isUser != "user" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// var JWTMiddlewareConfig = middleware.JWTWithConfig(middleware.JWTConfig{
// 	SigningMethod: "HS256",
// 	SigningKey:    []byte(constants.SECRET_KEY),
// 	TokenLookup:   "cookie:JWTCookie",
// 	AuthScheme:    "user",
// })
