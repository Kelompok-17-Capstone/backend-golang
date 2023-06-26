package payload

type UpdateOrder struct {
	Status    string `json:"status" form:"status"`
	ArrivedAt string `json:"arrived_at" form:"arrived_at"`
}
