package payload

type BalanceRequest struct {
	Total int `json:"total" form:"total" validate:"required,gte=0"`
}
