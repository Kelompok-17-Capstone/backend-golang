package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

// get all user

func GetUsers() ([]models.User, error) {
	var user []models.User
	if err := config.DB.Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// get user by id

func GetUser(id uint) (resp models.User, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}
