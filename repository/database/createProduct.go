package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func CreateProduct(req *models.Product) error {
	if err := config.DB.Save(&req).Error; err != nil {
		return err
	}
	return nil
}
