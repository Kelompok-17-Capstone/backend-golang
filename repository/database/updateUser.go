package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func UpdateUser(user *models.User) error {
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
