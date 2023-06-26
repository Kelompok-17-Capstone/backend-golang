package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetBalanceByUserID(userID uint) (*models.Balance, error) {
	resp := &models.Balance{}
	if err := config.DB.Where("user_id = ?", userID).First(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
