package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Order struct {
	Base
	UserID          uint           `json:"user_id" form:"user_id"`
	User            User           `json:"user" form:"user"`
	Address         string         `json:"address" form:"address"`
	BaseTotalPrice  int            `json:"base_total_price" form:"base_total_price"`
	Discount        int            `json:"discount" form:"discount" gorm:"default:0"`
	CoinsUsed       int            `json:"coins_used" form:"coins_used"`
	GrandTotalPrice int            `json:"grand_total_price" form:"grand_total_price"`
	Status          string         `json:"status" form:"status" gorm:"enum('dikemas','dikirim','diterima')"`
	ArrivedAt       datatypes.Date `json:"arrived_at" form:"arrived_at"  gorm:"default:null"`
	OrderDetails    []OrderDetail  `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

type OrderDetail struct {
	gorm.Model
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	OrderID   uuid.UUID `json:"order_id" form:"order_id"`
	Quantity  int       `json:"quantity" form:"quantity"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

func (o *Order) AfterDelete(tx *gorm.DB) (err error) {
	tx.Where("order_id = ?", o.ID).Delete(&OrderDetail{})
	return
}
