package payload

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GetOrderMobileParam struct {
	Status string
}

type GetOrderMobileData struct {
	OrderID         uuid.UUID             `json:"id" form:"order_id"`
	UserID          uint                  `json:"user_id" form:"user_id"`
	Address         string                `json:"address" form:"address"`
	BaseTotalPrice  int                   `json:"base_total_price" form:"base_total_price"`
	Discount        uint                  `json:"discount" form:"discount"`
	CoinUsed        uint                  `json:"coin_used" form:"coin_used"`
	GrandTotalPrice int                   `json:"grand_total_price" form:"grand_total_price"`
	Status          string                `json:"status" form:"status"`
	OrderAt         time.Time             `json:"order_at" form:"order_at"`
	ArrivedAt       time.Time             `json:"arrived_at" form:"arrived_at"`
	OrderDetails    []OrderDetailEmbedded `json:"order_details" form:"order_details"`
}

type OrderDetailEmbedded struct {
	OrderDetailID uint         `json:"id" form:"order_detail_id"`
	Quantity      int          `json:"quantity" form:"quantity"`
	Product       ProductEmbed `json:"product"`
}
