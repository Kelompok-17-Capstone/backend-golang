package payload

type UpdateOrderStatus struct {
	Status string `json:"status" form:"status" validate:"required"`
}

type UpadateArrivedAt struct {
	ArrivedAt string `json:"arrived_at" form:"arrived_at" validate:"required"`
}
