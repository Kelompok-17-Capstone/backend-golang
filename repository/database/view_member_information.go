package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

func ViewMemberInformation(id uint) (models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).Preload("Profile").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GenerateMemberCode() string {
	uuidObj := uuid.NewV4()
	return uuidObj.String()
}
