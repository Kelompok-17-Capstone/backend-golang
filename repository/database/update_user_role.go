package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func UpdateUserRole(id uint, role string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("role", role).Error; err != nil {
		return err
	}
	return nil
}

func GetUserProfile(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
