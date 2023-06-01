package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FavoriteProduct struct {
	gorm.Model
	ID        uint      `json:"id" form:"id"`
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	UserID    uint      `json:"user_id" form:"user_id"`
	Name      string    `json:"name" form:"name"`
	Price     uint      `json:"price" form:"price"`
	Image     string    `json:"image" form:"image"`
}
