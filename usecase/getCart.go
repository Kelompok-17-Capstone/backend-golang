package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
)

func GetCart(id uint) (resp payload.GetCart, err error) {
	cart, err := database.GetCart(id)
	if err != nil {
		return payload.GetCart{}, err
	}

	resp = payload.GetCart{
		ID:     cart.ID,
		UserID: cart.UserID,
	}
	for _, detail := range cart.DetailCartItems {
		item := payload.DetailCartItem{
			ID:         detail.ID,
			CartItemID: detail.CartItemID,
			Products: payload.ProductEmbed{
				ProductID: detail.ProductID,
				Name:      detail.Products.Name,
				Price:     detail.Products.Price,
				Image:     detail.Products.Image,
			},
			Quantity: detail.Quantity,
		}
		resp.DetailCartItem = append(resp.DetailCartItem, item)
	}
	return
}
