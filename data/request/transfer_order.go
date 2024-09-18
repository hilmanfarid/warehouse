package request

import (
	"golang-warehouse/model"
	"time"
)

type CreateTransferOrderRequest struct {
	ProductID            uint32                    `validate:"required" json:"product_id"`
	UserID               uint32                    `validate:"required" json:"user_id"`
	SourceWarehouse      uint32                    `validate:"required" json:"source_warehouse"`
	DestinationWarehouse uint32                    `validate:"required" json:"destination_warehouse"`
	Quantity             int32                     `validate:"required,min=1,max=100" json:"quantity"`
	Status               model.TransferOrderStatus `json:"status"`
	CreatedAt            time.Time                 `json:"created_at"`
	UpdatedAt            time.Time                 `json:"updated_at"`
	SuccessAt            time.Time                 `json:"success_at"`
	RefundedAt           time.Time                 `json:"refunded_at"`
	FailedAt             time.Time                 `json:"failed_at"`
}

type UpdateTransferOrderRequest struct {
	ID                   uint32                    `validate:"required"`
	ProductID            uint32                    `validate:"required" json:"product_id"`
	UserID               uint32                    `validate:"required" json:"user_id"`
	SourceWarehouse      uint32                    `validate:"required" json:"source_warehouse"`
	DestinationWarehouse uint32                    `validate:"required" json:"destination_warehouse"`
	Quantity             int32                     `validate:"required,min=1,max=100" json:"quantity"`
	Status               model.TransferOrderStatus `json:"status"`
	CreatedAt            time.Time                 `json:"created_at"`
	UpdatedAt            time.Time                 `json:"updated_at"`
	SuccessAt            time.Time                 `json:"success_at"`
	RefundedAt           time.Time                 `json:"refunded_at"`
	FailedAt             time.Time                 `json:"failed_at"`
}
