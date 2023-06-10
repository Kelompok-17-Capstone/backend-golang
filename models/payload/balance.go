package payload

type BalanceRequest struct {
	Total uint `json:"total" form:"total" validate:"required,gte=0"`
}
