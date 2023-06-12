package payload

type CreateAddress struct {
	ProfileId uint   `json:"profile_id" form:"profile_id"`
	Address   string `json:"address" form:"address" validate:"required"`
	City      string `json:"city" form:"city" validate:"required"`
	Province  string `json:"province" form:"province" validate:"required"`
}
