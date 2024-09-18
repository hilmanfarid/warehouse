package request

import (
	"golang-warehouse/model"
	"time"
)

type CreatePurchaseOrderDetailsRequest struct {
	PurchaseOrderID uint32                  `validate:"required,min=1,max=100" json:"name"`
	ProductID       uint32                  `validate:"required,min=1,max=100" json:"product_id"`
	WarehouseID     uint32                  `validate:"required,min=1,max=100" json:"warehouse_id"`
	Quantity        int32                   `validate:"required,min=1,max=100" json:"quantity"`
	PricePerUnit    int32                   `validate:"required,min=1,max=100" json:"price_per_unit"`
	Status          model.OrderDetailStatus `json:"status"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	SuccessAt       time.Time               `json:"success_at"`
	RefundedAt      time.Time               `json:"refunded_at"`
	FailedAt        time.Time               `json:"failed_at"`
}

type UpdatePurchaseOrderDetailsRequest struct {
	ID              uint32                  `validate:"required"`
	PurchaseOrderID uint32                  `validate:"required,min=1,max=100" json:"name"`
	ProductID       uint32                  `validate:"required,min=1,max=100" json:"product_id"`
	WarehouseID     uint32                  `validate:"required,min=1,max=100" json:"warehouse_id"`
	Quantity        int32                   `validate:"required,min=1,max=100" json:"quantity"`
	PricePerUnit    int32                   `validate:"required,min=1,max=100" json:"price_per_unit"`
	Status          model.OrderDetailStatus `json:"status"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	SuccessAt       time.Time               `json:"success_at"`
	FailedAt        time.Time               `json:"failed_at"`
	RefundedAt      time.Time               `json:"refunded_at"`
}
