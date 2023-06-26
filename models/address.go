package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ProfileID uint   `json:"profile_id" form:"profile_id"`
	Address   string `json:"address" form:"address"`
	Province  string `json:"province" form:"province"`
	City      string `json:"city" form:"city"`
	Status    bool   `json:"status" form:"status"`
}
