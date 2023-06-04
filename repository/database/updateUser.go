package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func UpdateUser(user *models.User, id uint) error {
	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}
	return nil
}
