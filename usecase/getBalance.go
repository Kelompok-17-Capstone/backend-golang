package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetBalance(userID uint) ([]payload.BalanceRequest, error) {
	balance, err := database.GetBalances(userID)
	if err != nil {
		return nil, err
	}

	var resp []payload.BalanceRequest
	for _, balance := range balance {
		resp = append(resp, payload.BalanceRequest{
			Total: balance.Total,
		})

		balance.Total = resp[len(resp)-1].Total
		if err := database.UpdateBalanceTotal(balance); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
