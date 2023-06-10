package models

import "gorm.io/gorm"

type Coin struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	BalanceID uint   `json:"balance_id"`
	Total     uint   `json:"total"`
	Status    string `json:"status" gorm:"type:enum('increase','decrease');default:'decrease'"`
}
