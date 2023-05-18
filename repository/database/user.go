package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetUserByEmail(email string) (user models.User, err error) {
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}
