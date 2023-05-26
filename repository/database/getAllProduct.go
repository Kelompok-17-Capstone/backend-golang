package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

// get all product

func GetAllProduct() ([]models.Product, error) {
	var product []models.Product
	if err := config.DB.Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

// get product by id

func GetProductById(id uuid.UUID) (resp models.Product, err error) {
	if err := config.DB.Where("id = ?", id).First(&resp).Error; err != nil {
		return resp, err
	}
	return
}
