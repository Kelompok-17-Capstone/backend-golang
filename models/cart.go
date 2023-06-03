package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UserID uint ` json:"user_id"`
	// TotalPrice     uint             ` json:"total_price"`
	DetailCartItems []DetailCartItem ` json:"detail_cart_item"`
}

type DetailCartItem struct {
	gorm.Model
	CartItemID uint      `json:"cart_item_id"`
	ProductID  uuid.UUID ` json:"product_id"`
	Quantity   uint      ` json:"quantity" gorm:"default:1"`
	Products   Product   `gorm:"foreignkey:ProductID"`
}

// User    User    `gorm:"foreignkey:UserID" json:"user"`
// Product Product `gorm:"foreignkey:ProductID" json:"product"`
