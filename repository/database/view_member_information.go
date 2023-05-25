package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func ViewMemberInformation(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
