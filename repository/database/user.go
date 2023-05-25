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

func CreateUserProfil(req *models.Profile) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserAddress(req *models.Address) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserStatus(id uint, status string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func UploadPhotoProfil(id uint, photo string) error {
	var profile *models.Profile
	if err := config.DB.Model(&profile).Where("user_id = ?", id).Update("photo", photo).Error; err != nil {
		return err
	}
	return nil
}
