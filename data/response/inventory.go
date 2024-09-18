package response

import "time"

type InventoryResponse struct {
	ID          uint32    `json:"id"`
	ProductID   uint32    `json:"product_id"`
	WarehouseID uint32    `json:"warehouse_id"`
	Quantity    int32     `json:"quantity"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
