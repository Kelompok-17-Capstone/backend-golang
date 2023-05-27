package payload

type GetUser struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required, min=11,max=13"`
	Address     string `json:"address" form:"address" validate:"required"`
}
