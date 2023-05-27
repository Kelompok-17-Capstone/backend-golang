package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

func UpdateProduct(id uuid.UUID, product *models.Product) error {
	if err := config.DB.Model(&product).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}
