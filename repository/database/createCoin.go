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

func UpdateCoin(id uint, total uint) error {
	var coin = models.Coin{}
	if err := config.DB.Model(&coin).Where("id = ?", id).Update("total", total).Error; err != nil {
		return err
	}
	return nil
}
