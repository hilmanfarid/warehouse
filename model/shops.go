package model

import (
	"golang-warehouse/data/response"
	"time"
)

const TableNameShop = "shops"

// Shop mapped from table <shops>
type Shop struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Status    int8      `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Shop's table name
func (*Shop) TableName() string {
	return TableNameShop
}

func (s Shop) ToResponse() (shop response.ShopResponse) {
	shop.ID = s.ID
	shop.Name = s.Name
	shop.Status = s.Status
	shop.CreatedAt = s.CreatedAt
	shop.UpdatedAt = s.UpdatedAt

	return shop
}
