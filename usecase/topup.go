package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"errors"

	"gorm.io/gorm"
)

func GetUserRole(userID uint) (string, error) {
	user, err := database.GetUserRole(userID)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func findCoinByID(topup *models.Balance, coinID uint) (*models.Coin, error) {
	for i := range topup.Coins {
		if topup.Coins[i].UserID == coinID {
			return &topup.Coins[i], nil
		}
	}
	return nil, nil
}
func CreateTopup(userID uint, total uint) error {
	existingTopup, err := database.GetBalanceByUserID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	userRole, err := GetUserRole(userID)
	if err != nil {
		return err
	}
	if userRole == "member" {
		if existingTopup == nil {
			// Saldo belum ada, buat saldo baru
			newTopup := &models.Balance{
				UserID: userID,
				Total:  total,
			}
			err = database.CreateTopup(newTopup)
			if err != nil {
				return err
			}
			coinTotal := uint(float32(total) * 0.01)
			newCoin := &models.Coin{
				UserID:    userID,
				BalanceID: newTopup.ID,
				Total:     coinTotal,
				Status:    "increase",
			}
			err = database.CreateCoin(newCoin)
			if err != nil {
				return err
			}
		} else {
			coin, err := findCoinByID(existingTopup, userID)
			if err != nil {
				return err
			}
			coinTotal := uint(float32(total) * 0.01)
			if coin != nil {
				// Pengguna sudah memiliki koin, perbarui nilai total koin
				coin.Total += coinTotal
				err = database.UpdateCoin(coin.ID, coin.Total)
				if err != nil {
					return err
				}
			}
			existingTopup.Total += total
			err = database.UpdateTopup(existingTopup.Total, existingTopup.ID)
			if err != nil {
				return err
			}
		}
	} else { // role != member
		if existingTopup == nil {
			// Saldo belum ada, buat saldo baru
			newTopup := &models.Balance{
				UserID: userID,
				Total:  total,
			}
			err = database.CreateTopup(newTopup)
			if err != nil {
				return err
			}
		} else {
			// Saldo sudah ada, perbarui saldo yang ada
			existingTopup.Total += total
			err = database.UpdateTopup(existingTopup.Total, existingTopup.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
