package response

import (
	"time"
)

type PurchaseOrderResponse struct {
	ID           uint32                        `json:"id"`
	Name         string                        `json:"name"`
	UserID       uint32                        `json:"user_id"`
	ShopID       uint32                        `json:"shop_id"`
	Status       string                        `json:"status"`
	TotalAmount  float64                       `json:"total_amount"`
	CreatedAt    time.Time                     `json:"created_at"`
	UpdatedAt    time.Time                     `json:"updated_at"`
	ProcessAt    time.Time                     `json:"process_at"`
	SuccessAt    time.Time                     `json:"success_at"`
	FailedAt     time.Time                     `json:"failed_at"`
	OrderDetails []PurchaseOrderDetailResponse `json:"order_details"`
}
