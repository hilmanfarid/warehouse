package model

import (
	"golang-warehouse/data/response"
	"time"
)

type TransferOrderStatus int8

const (
	TransferOrderStatusCreated TransferOrderStatus = iota + 1
	TransferOrderStatusSucceeded
	TransferOrderStatusFailed
	TransferOrderStatusRefunded
)

var transferOdcerStatusMap = map[TransferOrderStatus]string{
	TransferOrderStatusCreated:   "created",
	TransferOrderStatusSucceeded: "success",
	TransferOrderStatusFailed:    "failed",
	TransferOrderStatusRefunded:  "refunded",
}

func (s TransferOrderStatus) String() string {
	return transferOdcerStatusMap[s]
}

const TableNameTransferOrder = "transfer_orders"

// TransferOrder mapped from table <purchase_order_details>
type TransferOrder struct {
	ID                   uint32              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID               uint32              `gorm:"column:user_id" json:"user_id"`
	ProductID            uint32              `gorm:"column:product_id" json:"product_id"`
	SourceWarehouse      uint32              `gorm:"column:source_warehouse" json:"source_warehouse"`
	DestinationWarehouse uint32              `gorm:"column:destination_warehouse" json:"destination_warehouse"`
	Status               TransferOrderStatus `gorm:"column:status;not null" json:"status"`
	Quantity             int32               `gorm:"column:quantity;not null" json:"quantity"`
	CreatedAt            time.Time           `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt            time.Time           `gorm:"column:updated_at;not null" json:"updated_at"`
	SuccessAt            time.Time           `gorm:"column:success_at;default:null" json:"success_at"`
	FailedAt             time.Time           `gorm:"column:failed_at;default:null" json:"failed_at"`
}

// TableName TransferOrder's table name
func (*TransferOrder) TableName() string {
	return TableNameTransferOrder
}

func (pd TransferOrder) ToResponse() (pdResponse response.TransferOrderResponse) {
	pdResponse.ID = pd.ID
	pdResponse.UserID = pd.UserID
	pdResponse.ProductID = pd.ProductID
	pdResponse.SourceWarehouse = pd.SourceWarehouse
	pdResponse.DestinationWarehouse = pd.DestinationWarehouse
	pdResponse.Status = pd.Status.String()
	pdResponse.Quantity = pd.Quantity
	pdResponse.CreatedAt = pd.CreatedAt
	pdResponse.UpdatedAt = pd.UpdatedAt
	pdResponse.SuccessAt = pd.SuccessAt
	pdResponse.FailedAt = pd.FailedAt

	return pdResponse
}
