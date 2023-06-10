package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID      uint      `json:"user_id" form:"user_id" gorm:"unique"`
	Name        string    `json:"name" form:"name"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Address     []Address `gorm:"foreignKey:ProfileID"`
	Photo       string    `json:"photo" form:"photo"`
}
