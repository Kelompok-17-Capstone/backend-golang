package payload

type BalanceRequest struct {
	Total uint `json:"total" form:"total" svalidate:"required,min=11,max=13,gte=0"`
}
