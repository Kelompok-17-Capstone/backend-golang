package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func AddToCartItem(req *models.CartItem) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}

func AddToDetailCart(cartProduct *models.DetailCartItem) error {
	err := config.DB.Save(cartProduct).Error
	if err != nil {
		return err
	}
	return nil
}
func GetCartItemByUserID(id uint) (*models.CartItem, error) {
	resp := &models.CartItem{}
	if err := config.DB.Model(&models.CartItem{}).Preload("DetailCartItem.Product").Where("user_id = ?", id).First(resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func UpdateDetailCartItem(cartItem *models.DetailCartItem) error {
	if err := config.DB.Save(cartItem).Error; err != nil {
		return err
	}
	return nil
}
