package usecase

import (
	"backend-golang/models"
	"backend-golang/models/payload"
	"backend-golang/repository/database"

	"golang.org/x/crypto/bcrypt"
)

func Register(req *payload.Register) error {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    req.Email,
		Password: string(password),
	}
	if err := database.Register(&user); err != nil {
		return err
	}
	return nil
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
