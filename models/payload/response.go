package payload

import (
	uuid "github.com/satori/go.uuid"
)

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
type GetMemberMobile struct {
	ID          uint      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	Address     []Address `json:"address" form:"address"`
	Image       string    `json:"image" form:"image"`
	MemberCode  string    `json:"member_code" form:"member_code"`
	Barcode     string    `json:"barcode" form:"barcode"`
	Balance     int       `json:"balance" form:"balance"`
	Coin        int       `json:"coin" form:"coin"`
}

type Address struct {
	ID      uint   `json:"id" form:"id"`
	Address string `json:"address" form:"address"`
	Status  bool   `json:"status" form:"status"`
}
type GetCart struct {
	ID             uint             `json:"id" form:"id"`
	UserID         uint             `json:"user_id" form:"user_id"`
	DetailCartItem []DetailCartItem `json:"detail_cart_item"`
}
type ProductEmbed struct {
	ProductID uuid.UUID `json:"id" form:"product_id"`
	Name      string    `json:"name" form:"name"`
	Price     uint      `json:"price" form:"price"`
	Image     string    `json:"image" form:"image"`
}
type DetailCartItem struct {
	ID         uint `json:"id" form:"id"`
	CartItemID uint `json:"cart_item_id"`
	// ProductID  uuid.UUID      `json:"product_id"`
	Quantity uint         `json:"quantity" validate:"gt=0"`
	Products ProductEmbed `json:"product"`
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
