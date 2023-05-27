package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetProductsByNameAndStock(name string, sign string) ([]models.Product, error) {
	var products []models.Product

	if err := config.DB.Where("name like ? AND stock "+sign+" ?", "%"+name+"%", 0).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
