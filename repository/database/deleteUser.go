package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func DeleteUser(id uint) error {
	user := models.User{}
	if err := config.DB.Model(&user).Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
