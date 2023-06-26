package models

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Total  int    `json:"total" form:"total"`
	Status string `json:"status" gorm:"type:enum('increase','decrease')"`
}
