package payload

import (
	"time"
)

type GetNotification struct {
	ID        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	Text      string    `json:"text" form:"text"`
	IsRead    bool      `json:"is_read" form:"is_read"`
}
