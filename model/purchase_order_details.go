package model

import (
	"golang-warehouse/data/response"
	"time"
)

type OrderDetailStatus int8

const (
	OrderDetailStatusCreated OrderDetailStatus = iota + 1
	OrderDetailStatusSucceeded
	OrderDetailStatusFailed
	OrderDetailStatusRefunded
)

var orderDetailStatusMap = map[OrderDetailStatus]string{
	OrderDetailStatusCreated:   "created",
	OrderDetailStatusSucceeded: "success",
	OrderDetailStatusFailed:    "failed",
	OrderDetailStatusRefunded:  "refunded",
}

func (s OrderDetailStatus) String() string {
	return orderDetailStatusMap[s]
}

const ()

const TableNamePurchaseOrderDetail = "purchase_order_details"

// PurchaseOrderDetail mapped from table <purchase_order_details>
type PurchaseOrderDetail struct {
	ID              uint32            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PurchaseOrderID uint32            `gorm:"column:purchase_order_id" json:"purchase_order_id"`
	ProductID       uint32            `gorm:"column:product_id" json:"product_id"`
	WarehouseID     uint32            `gorm:"column:warehouse_id;default:null" json:"warehouse_id"`
	Status          OrderDetailStatus `gorm:"column:status;not null" json:"status"`
	Quantity        int32             `gorm:"column:quantity;not null" json:"quantity"`
	PricePerUnit    float64           `gorm:"column:price_per_unit" json:"price_per_unit"`
	CreatedAt       time.Time         `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"column:updated_at;not null" json:"updated_at"`
	SuccessAt       time.Time         `gorm:"column:success_at;default:null" json:"success_at"`
	FailedAt        time.Time         `gorm:"column:failed_at;default:null" json:"failed_at"`
	RefundedAt      time.Time         `gorm:"column:refunded_at;default:null" json:"refunded_at"`
}

// TableName PurchaseOrderDetail's table name
func (*PurchaseOrderDetail) TableName() string {
	return TableNamePurchaseOrderDetail
}

func (pd PurchaseOrderDetail) ToResponse() (pdResponse response.PurchaseOrderDetailResponse) {
	pdResponse.ID = pd.ID
	pdResponse.PurchaseOrderID = pd.PurchaseOrderID
	pdResponse.ProductID = pd.ProductID
	pdResponse.WarehouseID = pd.WarehouseID
	pdResponse.Status = pd.Status.String()
	pdResponse.Quantity = pd.Quantity
	pdResponse.PricePerUnit = pd.PricePerUnit
	pdResponse.CreatedAt = pd.CreatedAt
	pdResponse.UpdatedAt = pd.UpdatedAt
	pdResponse.SuccessAt = pd.SuccessAt
	pdResponse.FailedAt = pd.FailedAt
	pdResponse.RefundedAt = pd.RefundedAt

	return pdResponse
}
