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
	if err := config.DB.Model(&models.CartItem{}).Preload("DetailCartItems.Products").Where("user_id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return resp, nil
}

func UpdateDetailCartItem(quantity uint, cartItemID uint) error {
	if err := config.DB.Model(&models.DetailCartItem{}).Where("id = ?", cartItemID).Update("quantity", quantity).Error; err != nil {
		return err
	}
	return nil
}
