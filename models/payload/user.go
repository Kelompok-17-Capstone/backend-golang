package payload

type GetUser struct {
	ID          uint   `json:"id" format:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	Status      string `json:"status" from:"status"`
}
