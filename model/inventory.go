package model

import (
	"golang-warehouse/data/response"
	"time"
)

const TableNameInventory = "inventory"

// Inventory mapped from table <inventory>
type Inventory struct {
	ID          uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductID   uint32    `gorm:"column:product_id" json:"product_id"`
	WarehouseID uint32    `gorm:"column:warehouse_id" json:"warehouse_id"`
	Quantity    int32     `gorm:"column:quantity;not null" json:"quantity"`
	Status      int8      `gorm:"column:status;not null" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Inventory's table name
func (*Inventory) TableName() string {
	return TableNameInventory
}

func (s Inventory) ToResponse() (inventory response.InventoryResponse) {
	inventory.ID = s.ID
	inventory.ProductID = s.ProductID
	inventory.WarehouseID = s.WarehouseID
	inventory.Quantity = s.Quantity
	inventory.Status = s.Status
	inventory.CreatedAt = s.CreatedAt
	inventory.UpdatedAt = s.UpdatedAt

	return inventory
}
