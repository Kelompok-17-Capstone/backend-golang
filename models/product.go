package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	Base
	ProductName string `json:"product_name" form:"product_name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
	ImageFile   string `json:"-" gorm:"column:image_file"`
}
type Base struct {
	// ID        uuid.UUID  `gorm:"primary_key" json:"id"`
	ID        uuid.UUID  `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
// 	if base.ID == uuid.Nil {
// 		base.ID = uuid.NewV4()
// 	}
// 	return nil
// }

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
