package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}
