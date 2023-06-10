package usecase

import (
	"backend-golang/models"
	"backend-golang/repository/database"
)

func CreateTopup(userID uint, total uint) error {
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

	if user.Role == "member" {
		coinTotal := uint(float32(total) * 0.01)
		coin := &models.Coin{
			UserID: userID,
			Total:  coinTotal,
			Status: "increase",
		}
		if err := database.CreateCoin(coin); err != nil {
			return err
		}
		user.Coin += coinTotal
	}

	if err := database.UpdateUser(&user); err != nil {
		return err
	}
	return nil
}
