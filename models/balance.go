package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Total  uint `json:"total"`
}
