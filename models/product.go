package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
	Image       string `json:"image" form:"image"`
}
type Base struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
