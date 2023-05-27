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
<<<<<<< HEAD

type UpdateProductRespone struct {
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       uint      `json:"stock" form:"stock"`
	Price       uint      `json:"price" form:"price"`
	Image       string    `json:"image" form:"image"`
=======
type GetMember struct {
	ID         uint   `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Image      string `json:"image" form:"image"`
	MemberCode string `json:"member_code" form:"member_code"`
>>>>>>> 51b84e3bb44798f2f6db55fff0d9db64e1015c35
}
