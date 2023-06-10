package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func CreateCoin(coin *models.Coin) error {
	if err := config.DB.Save(&coin).Error; err != nil {
		return err
	}
	return nil
}
