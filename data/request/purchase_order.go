package request

import (
	"golang-warehouse/model"
	"time"
)

type CreatePurchaseOrderRequest struct {
	UserID       uint32            `json:"user_id"`
	ShopID       uint32            `validate:"required,min=1,max=100" json:"shop_id"`
	Status       model.OrderStatus `form:"default=1" json:"status"`
	OrderDetails []struct {
		ProductID uint32 `validate:"required,min=1,max=100" json:"product_id"`
		Quantity  int32  `validate:"required,min=1,max=100" json:"quantity"`
	} `validate:"required" json:"order_details"`
}

type UpdatePurchaseOrderRequest struct {
	ID          uint32            `validate:"required"`
	UserID      uint32            `validate:"required,min=1,max=100" json:"user_id"`
	ShopID      uint32            `validate:"required,min=1,max=100" json:"shop_id"`
	Status      model.OrderStatus `json:"status"`
	TotalAmount float64           `json:"total_amount"`
	UpdatedAt   time.Time         `json:"updated_at"`
	ProcessAt   time.Time         `json:"process_at"`
	SuccessAt   time.Time         `json:"success_at"`
	FailedAt    time.Time         `json:"failed_at"`
}
