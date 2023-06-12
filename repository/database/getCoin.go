package database

import (
	"backend-golang/config"
	"backend-golang/models"
)

func GetCoin(userID uint) ([]*models.Coin, error) {
	var resp []*models.Coin
	if err := config.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

// func UpdateCoinTotal(coin *models.Coin) error {
// 	if err := config.DB.Save(coin).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
