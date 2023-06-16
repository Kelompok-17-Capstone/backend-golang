package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetCart(userID uint) (*models.CartItem, error) {
	resp := &models.CartItem{}
	if err := config.DB.Preload("DetailCartItems.Products").Where("user_id = ?", userID).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}
