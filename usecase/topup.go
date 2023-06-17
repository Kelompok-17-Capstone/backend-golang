package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
	"strconv"
)

func CreateTopup(userID uint, total int) error {
	user, err := database.GetUser(userID)
	if err != nil {
		return err
	}
	topUp := &models.Balance{
		UserID: userID,
		Total:  total,
		Status: "increase",
	}

	if err := database.CreateTopup(topUp); err != nil {
		return err
	}

	user.Balance += total
	balance := strconv.Itoa(total)

	if user.Role == "member" {
		coinTotal := int(float32(total) * 0.01)
		coin := &models.Coin{
			UserID: userID,
			Total:  coinTotal,
			Status: "increase",
		}
		if err := database.CreateCoin(coin); err != nil {
			return err
		}
		user.Coin += coinTotal

		coins := strconv.Itoa(coinTotal)

		if err := database.SaveNotification(models.Notification{
			UserID: userID,
			Text:   "Selamat anda mendapatkan bonus koin sebesar " + coins + ", karena anda telah melakukan top up saldo sebesar " + balance,
		}); err != nil {
			return err
		}
	} else {
		if err := database.SaveNotification(models.Notification{
			UserID: userID,
			Text:   "Selamat anda telah berhasil melakukan top up sebesar " + balance,
		}); err != nil {
			return err
		}
	}

	if err := database.UpdateUser(&user); err != nil {
		return err
	}
	return nil
}
