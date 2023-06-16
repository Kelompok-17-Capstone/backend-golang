package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Text   string `json:"text" form:"text"`
	IsRead bool   `json:"is_read" form:"is_read" gorm:"default:false"`
}
