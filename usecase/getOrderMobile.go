package usecase

import (
	"backend-golang/models/payload"
	"backend-golang/repository/database"
	"time"
)

func GetOrderMobile(userID uint, status string) ([]payload.GetOrderMobileData, error) {
	req := payload.GetOrderMobileParam{
		Status: status,
	}
	user, err := database.GetUser(userID)
	if err != nil {
		return nil, err
	}
	orderDetails, err := database.GetOrderMobile(userID, &req)
	if err != nil {
		return nil, err
	}
	orderResponses := make([]payload.GetOrderMobileData, 0, len(orderDetails))
	for _, order := range orderDetails {
		var orderDetailsPayload []payload.OrderDetailEmbedded

		for _, detail := range order.OrderDetails {
			orderDetailsPayload = append(orderDetailsPayload, payload.OrderDetailEmbedded{
				OrderDetailID: detail.ID,
				Quantity:      detail.Quantity,
				Product: payload.ProductEmbed{
					ProductID: detail.ProductID,
					Name:      detail.Product.Name,
					Price:     detail.Product.Price,
					Image:     detail.Product.Image,
				},
			})
		}
		orderResponse := payload.GetOrderMobileData{
			UserID:          user.ID,
			OrderID:         order.ID,
			Address:         order.Address,
			BaseTotalPrice:  order.BaseTotalPrice,
			CoinUsed:        uint(order.CoinsUsed),
			Discount:        uint(order.Discount),
			GrandTotalPrice: order.GrandTotalPrice,
			Status:          order.Status,
			OrderAt:         time.Time(order.CreatedAt),
			ArrivedAt:       time.Time(order.ArrivedAt),
			OrderDetails:    orderDetailsPayload,
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	return orderResponses, err
}
