package payload

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
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
}

type UpdateProduct struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Stock       uint   `json:"stock" form:"stock"`
	Price       uint   `json:"price" form:"price"`
	Image       string `json:"image" form:"image"`
}
