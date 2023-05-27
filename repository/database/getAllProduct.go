package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

// get all product

func GetAllProduct() ([]models.Product, error) {
	var product []models.Product
	if err := config.DB.Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
