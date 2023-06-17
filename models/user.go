package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string            `json:"email" form:"email" gorm:"unique"`
	Password      string            `json:"password" form:"password"`
	Role          string            `json:"role" form:"role" gorm:"type:enum('reguler','member','admin');default:'reguler'"`
	Status        string            `json:"status" form:"status" gorm:"type:enum('validated','unvalidated');default:'unvalidated'"`
	MemberCode    string            `json:"member_code" form:"member_code"`
	Barcode       string            `json:"barcode" form:"barcode"`
	Orders        []Order           `gorm:"foreignKey:UserID"`
	Carts         CartItem          `gorm:"foreignKey:UserID"`
	Balance       int               `json:"balance" form:"balance"`
	Coin          int               `json:"coin" form:"coin"`
	Profile       Profile           `gorm:"foreignKey:UserID"`
	Favorites     []FavoriteProduct `gorm:"foreignKey:UserID"`
	Balances      []Balance         `gorm:"foreignKey:UserID"`
	Coins         []Coin            `gorm:"foreignKey:UserID"`
	Notifications []Notification    `gorm:"foreignKey:UserID"`
}

func (user *User) BeforeDelete(tx *gorm.DB) error {
	if user.Role == "admin" {
		return echo.NewHTTPError(http.StatusBadRequest, "Cant delete admin user")
	}
	return nil
}
