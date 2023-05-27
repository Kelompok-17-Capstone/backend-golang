package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string  `json:"email" form:"email" gorm:"unique"`
	Password   string  `json:"password" form:"password"`
	Role       string  `json:"role" form:"role" gorm:"type:enum('reguler','member','admin');default:'reguler'"`
	Status     string  `json:"status" form:"status" gorm:"type:enum('validated','unvalidated');default:'unvalidated'"`
	MemberCode string  `json:"member_code" form:"member_code"`
	Profile    Profile `gorm:"foreignKey:UserID"`
}
