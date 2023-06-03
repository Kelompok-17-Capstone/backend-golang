package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func UpdateCartQty(id uint, req *payload.UpdateQtyCart) error {
	detailCartItem, err := database.GetDetailCartItemById(id)
	if err != nil {
		return err
	}

	detailCartItem.Quantity = req.Quantity
	if err := database.UpdateDetailCartItem(detailCartItem.Quantity, id); err != nil {
		return err
	}

	return nil
}
