package database

import (
	"backend-golang/config"
	"backend-golang/models"
	"backend-golang/models/payload"
)

// get all user

func GetUsers(req *payload.UsersParam) ([]models.User, error) {
	var user []models.User
	db := config.DB

	if req.Keyword != "" {
		db = db.Joins("JOIN profiles ON users.id = profiles.user_id").Where("profiles.name like ?", "%"+req.Keyword+"%")
	}
	if req.Role != "" {
		db = db.Where("role = ?", req.Role)
	}
	if err := db.Preload("Profile.Address").Not("role = ?", "admin").Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// get user by id

func GetUser(id uint) (resp models.User, err error) {
	if err := config.DB.Preload("Profile.Address").Preload("Favorites.Product").Preload("Notifications").Where("id = ?", id).Not("role = ?", "admin").First(&resp).Error; err != nil {
		return resp, err
	}
	return
}
