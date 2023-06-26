package payload

import (
	"gorm.io/datatypes"
)

type DashboardResponse struct {
	Users    []UsersResponse    `json:"users"`
	Orders   []OrdersResponse   `json:"orders"`
	Products []ProductsResponse `json:"products"`
	Graphic  []Graphic          `json:"graphic"`
}

type UsersResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Role        string `json:"role" form:"role" gorm:"type:enum('reguler','member','admin');default:'reguler'"`
}

type OrdersResponse struct {
	UserName    string         `json:"user_name"`
	Address     string         `json:"address"`
	ProductName string         `json:"product_name"`
	Quantity    int            `json:"quantity"`
	OrderAt     datatypes.Date `json:"order_at"`
}

type ProductsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	Price       uint   `json:"price"`
	Status      string `json:"status" form:"status" gorm:"type:enum('Tersedia','Tidak Tersedia');default:'Tersedia'"`
}

type OrderDetailsResponse struct {
	ProductName string `json:"product"`
	Quantity    int    `json:"quantity"`
}

type Graphic struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}
