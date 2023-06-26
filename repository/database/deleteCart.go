package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

func DeleteDetailCartItem(id uint) error {
	if err := config.DB.Delete(&models.DetailCartItem{}, id).Error; err != nil {
		return err
	}
	return nil
}
func GetDetailCartItemById(id uint) (*models.DetailCartItem, error) {
	cart := &models.DetailCartItem{}
	if err := config.DB.Where("id = ?", id).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func DeleteCartByProductId(id uuid.UUID) error {
	var cart models.DetailCartItem
	if err := config.DB.Model(&cart).Where("product_id = ?", id).Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}
