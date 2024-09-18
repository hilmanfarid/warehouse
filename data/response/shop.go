package response

import (
	"time"
)

type ShopResponse struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
