package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func CreateTopup(topup *models.Balance) error {
	if err := config.DB.Save(&topup).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTopup(total uint, balanceID uint) error {
	if err := config.DB.Model(&models.Balance{}).Where("id = ?", balanceID).Update("total", total).Error; err != nil {
		return err
	}
	return nil
}
