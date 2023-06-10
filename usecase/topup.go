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
				UserID: userID,
				Total:  coinTotal,
				Status: "increase",
			}
			err = database.CreateCoin(newCoin)
			if err != nil {
				return err
			}
		} else {
			existingTopup.Total += total
			err = database.UpdateTopup(existingTopup.Total, existingTopup.ID)
			if err != nil {
				return err
			}
			// Saldo sudah ada, tidak perlu melakukan pembaruan pada koin
			coinTotal := uint(float32(total) * 0.01)
			newCoin := &models.Coin{
				UserID: userID,
				Total:  coinTotal,
				Status: "increase",
			}
			err = database.CreateCoin(newCoin)
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
