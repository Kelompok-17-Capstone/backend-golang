package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

// update user email
func UpdateUserEmail(id uint, email string) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", id).Update("email", email).Error; err != nil {
		return err
	}
	return nil
}

// update user password
func GetUserById(id uint) (user models.User, err error) {
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return
}

// upadate Name User
func UpdateName(id uint, name string) error {
	if err := config.DB.Model(&models.Profile{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

// update phone number
func UpdatePhoneNumber(id uint, phoneNumber string) error {
	if err := config.DB.Model(&models.Profile{}).Where("id = ?", id).Update("phone_number", phoneNumber).Error; err != nil {
		return err
	}
	return nil
}
