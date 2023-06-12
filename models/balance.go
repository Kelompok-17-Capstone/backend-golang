package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Total  int    `json:"total" form:"total"`
	Status string `json:"status" form:"status" gorm:"type:enum('increase','decrease')"`
}
