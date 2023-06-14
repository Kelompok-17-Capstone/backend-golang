package payload

type BalanceRequest struct {
	Total int `json:"total" form:"total" validate:"gt=0"`
}
