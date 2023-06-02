package payload

import uuid "github.com/satori/go.uuid"

type ProductResponse struct {
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       uint      `json:"stock" form:"stock"`
	Price       uint      `json:"price" form:"price"`
	Image       string    `json:"image" form:"image"`
}
type GetMember struct {
	ID         uint   `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Image      string `json:"image" form:"image"`
	MemberCode string `json:"member_code" form:"member_code"`
}
type GetCart struct {
	ID     uint `json:"id" form:"id"`
	UserID uint `json:"user_id" form:"user_id"`
	// ProductID      uuid.UUID        `json:"product_id" form:"product_id"`
	// Quantity       uint             `json:"quatity" form:"quantity"`
	DetailCartItem []DetailCartItem `json:"detail_cart_item"`
}
type DetailCartItem struct {
	CartItemID uint      `json:"cart_item_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   uint      `json:"quantity"`
}
