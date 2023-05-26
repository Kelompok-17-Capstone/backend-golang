package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetProductsByStock(sign string) ([]models.Product, error) {
	var products []models.Product

	if err := config.DB.Where("stock "+sign+" ?", 0).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
