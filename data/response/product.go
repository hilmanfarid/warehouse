package response

import "time"

type ProductResponse struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Status    int8      `json:"status"`
	Code      string    `json:"code"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
