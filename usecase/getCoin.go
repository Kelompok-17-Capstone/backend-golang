package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetCoin(userID uint) ([]payload.CoinResponse, error) {
	coins, err := database.GetCoin(userID)
	if err != nil {
		return nil, err
	}

	var resp []payload.CoinResponse
	for _, coin := range coins {
		resp = append(resp, payload.CoinResponse{
			Total:  coin.Total,
			Status: coin.Status,
		})

		coin.Total = resp[len(resp)-1].Total
		if err := database.UpdateCoinTotal(coin); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
