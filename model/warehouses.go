package model

import (
	"golang-warehouse/data/response"
	"time"
)

const TableNameWarehouse = "warehouses"

// Warehouse mapped from table <warehouses>
type Warehouse struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ShopID    uint32    `gorm:"column:shop_id" json:"shop_id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Status    int8      `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Warehouse's table name
func (*Warehouse) TableName() string {
	return TableNameWarehouse
}

func (s Warehouse) ToResponse() (warehouse response.WarehouseResponse) {
	warehouse.ID = s.ID
	warehouse.ShopID = s.ShopID
	warehouse.Name = s.Name
	warehouse.Status = s.Status
	warehouse.CreatedAt = s.CreatedAt
	warehouse.UpdatedAt = s.UpdatedAt

	return warehouse
}
