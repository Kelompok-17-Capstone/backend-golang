package payload

import uuid "github.com/satori/go.uuid"

type ProductResponse struct {
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       uint      `json:"stock" form:"stock"`
	Price       uint      `json:"price" form:"price"`
	Image       string    `json:"image" form:"image"`
	Status      string    `json:"status" form:"status"`
}
type ProductMobileResponse struct {
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       uint      `json:"stock" form:"stock"`
	Price       uint      `json:"price" form:"price"`
	Image       string    `json:"image" form:"image"`
	Status      string    `json:"status" form:"status"`
	Favorit     int       `json:"favorit" form:"favorit"`
}
type GetMember struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	Image       string `json:"image" form:"image"`
	MemberCode  string `json:"member_code" form:"member_code"`
}

type GetFavoriteProduct struct {
	ID          uint      `json:"id" form:"id"`
	ProductID   uuid.UUID `json:"product_id" form:"product_id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Price       uint      `json:"price" form:"price"`
	Image       string    `json:"image" form:"image"`
	Favorite    bool      `json:"favorite" form:"favorite"`
}
