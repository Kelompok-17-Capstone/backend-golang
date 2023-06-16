package payload

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type CreateOrder struct {
	Products []Product `json:"products" form:"products"`
	Address  string    `json:"address" form:"address"`
	Cart     bool      `json:"cart" form:"cart"`
	Coin     bool      `json:"coin" form:"coin"`
}

type Product struct {
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	Quantity  int       `json:"quantity" form:"quantity" validate:"gt=0"`
}

type GetOrders struct {
	ID            uuid.UUID      `json:"id" form:"id"`
	Name          string         `json:"name" form:"name"`
	Address       string         `json:"address" form:"address"`
	Products      string         `json:"product" form:"product"`
	TotalQuantity int            `json:"total_quantity" form:"total_quantity"`
	TotalPrice    int            `json:"total_price" form:"total_price"`
	OrderAt       datatypes.Date `json:"order_at" form:"order_at"`
	ArrivedAt     datatypes.Date `json:"arrived_at" form:"arrived_at"`
	Status        string         `json:"status" form:"status"`
}
