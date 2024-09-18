package model

import (
	"golang-warehouse/data/response"
	"time"
)

type OrderStatus int8

const (
	OrderStatusCreated OrderStatus = iota + 1
	OrderStatusProcessed
	OrderStatusSucceeded
	OrderStatusFailed
)

var orderStatusMap = map[OrderStatus]string{
	OrderStatusCreated:   "created",
	OrderStatusProcessed: "processed",
	OrderStatusSucceeded: "success",
	OrderStatusFailed:    "failed",
}

var orderStatusMapstring = map[string]OrderStatus{
	"created":   OrderStatusCreated,
	"processed": OrderStatusProcessed,
	"success":   OrderStatusSucceeded,
	"failed":    OrderStatusFailed,
}

func (s OrderStatus) String() string {
	return orderStatusMap[s]
}

func OderStatusFromString(val string) OrderStatus {
	return orderStatusMapstring[val]
}

const TableNamePurchaseOrder = "purchase_orders"

// PurchaseOrder mapped from table <purchase_orders>
type PurchaseOrder struct {
	ID          uint32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID      uint32      `gorm:"column:user_id" json:"user_id"`
	ShopID      uint32      `gorm:"column:shop_id" json:"shop_id"`
	Status      OrderStatus `gorm:"column:status;not null" json:"status"`
	TotalAmount float64     `gorm:"column:total_amount" json:"total_amount"`
	CreatedAt   time.Time   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;not null" json:"updated_at"`
	ProcessAt   time.Time   `gorm:"column:process_at;default:null" json:"process_at"`
	SuccessAt   time.Time   `gorm:"column:success_at;default:null" json:"success_at"`
	FailedAt    time.Time   `gorm:"column:failed_at;default:null" json:"failed_at"`
}

// TableName PurchaseOrder's table name
func (*PurchaseOrder) TableName() string {
	return TableNamePurchaseOrder
}

func (s PurchaseOrder) ToResponse() (purchaseOrder response.PurchaseOrderResponse) {
	purchaseOrder.ID = s.ID
	purchaseOrder.UserID = s.UserID
	purchaseOrder.ShopID = s.ShopID
	purchaseOrder.Status = s.Status.String()
	purchaseOrder.TotalAmount = s.TotalAmount
	purchaseOrder.CreatedAt = s.CreatedAt
	purchaseOrder.UpdatedAt = s.UpdatedAt
	purchaseOrder.ProcessAt = s.ProcessAt
	purchaseOrder.SuccessAt = s.SuccessAt
	purchaseOrder.FailedAt = s.FailedAt

	return purchaseOrder
}
