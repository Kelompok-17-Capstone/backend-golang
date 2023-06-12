package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetBalances(userID uint) ([]*models.Balance, error) {
	var resp []*models.Balance
	if err := config.DB.Where("user_id = ?", userID).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateBalanceTotal(balance *models.Balance) error {
	if err := config.DB.Save(balance).Error; err != nil {
		return err
	}
	return nil
}
