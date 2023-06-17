package database

import (
	"backend-golang/config"
	"backend-golang/models"

	uuid "github.com/satori/go.uuid"
)

func DeleteProduct(id uuid.UUID) error {
	product := models.Product{}
	if err := config.DB.Model(&product).Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
