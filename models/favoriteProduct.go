package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type FavoriteProduct struct {
	gorm.Model
	Product   Product   `gorm:"foreignKey:ProductID"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" gorm:"foreignKey:ProductID"`
	UserID    uint      `json:"user_id" form:"user_id" gorm:"foreignKey:UserID"`
}

func (fav *FavoriteProduct) BeforeDelete(tx *gorm.DB) error {
	var product Product
	tx.First(&product, fav.ProductID)
	product.Favorite--
	tx.Model(&product).Where("id = ?", fav.ProductID).Save(&product)
	return nil
}

func (fav *FavoriteProduct) AfterCreate(tx *gorm.DB) error {
	var product Product
	tx.First(&product, fav.ProductID)
	product.Favorite++
	tx.Model(&product).Where("id = ?", fav.ProductID).Updates(&product)
	return nil
}
