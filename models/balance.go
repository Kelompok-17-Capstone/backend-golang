package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Status string `json:"status" form:"status"`
	Total  uint   `json:"total" form:"total"`
}
