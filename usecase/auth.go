package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(req *payload.Register) (models.User, error) {
	if req.Password != req.RetypePassword {
		return models.User{}, echo.NewHTTPError(http.StatusBadRequest, "Password not match")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{
		Email:    req.Email,
		Password: string(password),
	}
	if err := database.Register(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func Login(req *payload.Login) (models.User, error) {
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return models.User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return models.User{}, err
	}
	return user, nil
}
