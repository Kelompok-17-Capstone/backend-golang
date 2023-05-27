package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetProductsByName(name string) ([]models.Product, error) {
	var products []models.Product

	if err := config.DB.Where("name like ?", "%"+name+"%").Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
