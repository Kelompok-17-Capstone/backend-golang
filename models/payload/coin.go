package payload

import "time"

type CoinResponse struct {
	Total     uint      `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
