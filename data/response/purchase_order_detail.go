package response

import "time"

type PurchaseOrderDetailResponse struct {
	ID              uint32    `json:"id"`
	PurchaseOrderID uint32    `json:"purchase_order_id"`
	ProductID       uint32    `json:"product_id"`
	WarehouseID     uint32    `json:"warehouse_id"`
	Status          string    `json:"status"`
	Quantity        int32     `json:"quantity"`
	PricePerUnit    float64   `json:"price_per_unit"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	SuccessAt       time.Time `json:"success_at"`
	FailedAt        time.Time `json:"failed_at"`
	RefundedAt      time.Time `json:"refunded_at"`
}
