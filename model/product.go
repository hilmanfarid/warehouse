package model

import (
	"golang-warehouse/data/response"
	"time"
)

const TableNameProduct = "products"

// Product mapped from table <products>
type Product struct {
	ID        uint32    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	Code      string    `gorm:"column:code;not null" json:"code"`
	Status    int8      `gorm:"column:status" json:"status"`
	Price     float64   `gorm:"column:price" json:"price"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName Product's table name
func (*Product) TableName() string {
	return TableNameProduct
}

func (s Product) ToResponse() (product response.ProductResponse) {
	product.ID = s.ID
	product.Name = s.Name
	product.Status = s.Status
	product.Code = s.Code
	product.Price = s.Price
	product.CreatedAt = s.CreatedAt
	product.UpdatedAt = s.UpdatedAt

	return product
}
