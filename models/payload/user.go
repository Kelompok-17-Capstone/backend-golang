package payload

type GetUser struct {
	ID          uint   `json:"id" format:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	Status      string `json:"status" from:"status"`
}

type UpdateUser struct {
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,min=11,max=13"`
	City        string `json:"city" form:"city" validate:"required"`
	Province    string `json:"province" form:"province" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Status      string `json:"status" form:"status" validate:"required"`
}
