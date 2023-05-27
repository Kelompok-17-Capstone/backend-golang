package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

// get all user

func GetUsers() ([]models.User, error) {
	var user []models.User
	if err := config.DB.Preload("Profile.Address").Not("role = ?", "admin").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// get user by id

func GetUser(id uint) (resp models.User, err error) {
	if err := config.DB.Preload("Profile.Address").Where("id = ?", id).Not("role = ?", "admin").First(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetUsersByName(name string) (resp []models.User, err error) {
	if err := config.DB.Joins("JOIN profiles ON users.id = profiles.user_id").Where("profiles.name like ?", "%"+name+"%").Preload("Profile.Address").Not("role = ?", "admin").Find(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetUsersByRole(role string) (resp []models.User, err error) {
	if err := config.DB.Preload("Profile.Address").Where("role = ?", role).Find(&resp).Error; err != nil {
		return resp, err
	}
	return
}

func GetUsersByNameAndRole(name, role string) (resp []models.User, err error) {
	if err := config.DB.Joins("JOIN profiles ON users.id = profiles.user_id").Where("profiles.name like ?", "%"+name+"%").Preload("Profile.Address").Where("role = ?", role).Find(&resp).Error; err != nil {
		return resp, err
	}
	return
}
