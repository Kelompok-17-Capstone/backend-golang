package payload

import uuid "github.com/satori/go.uuid"

type Register struct {
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=8"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type ProductParam struct {
	Keyword string
	Status  string
	Price   string
	Tab     string
}

type OrdersParam struct {
	Keyword string
	Status  string
}

type UsersParam struct {
	Keyword string
	Role    string
}

type Profile struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,min=11,max=13"`
	City        string `json:"city" form:"city" validate:"required"`
	Province    string `json:"province" form:"province" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
}

type CreateProduct struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock" validate:"gt=0"`
	Price       uint   `json:"price" form:"price"`
}

type UpdateProduct struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock" validate:"gte=0"`
	Price       uint   `json:"price" form:"price"`
	Image       string `json:"image" form:"image"`
}

type UpdateEmail struct {
	Email string `json:"email" form:"email" validate:"email"`
}

type UpdatePhoneNumber struct {
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"min=11,max=13"`
}
type UpdatePassword struct {
	OldPassword    string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword    string `json:"new_password" form:"new_password" validate:"required,min=8"`
	RetypePassword string `json:"retype_password" form:"retype_password" validate:"required"`
}

type AddFavoriteProduct struct {
	ProductID string `json:"product_id" form:"product_id"`
}

type UpdateAddress struct {
	Address  string `json:"address" form:"address"`
	Province string `json:"province" form:"province"`
	City     string `json:"city" form:"city"`
	Status   bool   `json:"status" form:"status"`
}

type AddToCart struct {
	UserID    uint      `json:"user_id" form:"user_id"`
	ProductID uuid.UUID `json:"product_id" form:"product_id"`
	Quantity  uint      `json:"quantity" form:"quantity" validate:"gt=0"`
}
type UpdateQtyCart struct {
	Quantity uint `json:"quantity" form:"quantity" validate:"gte=0"`
}
