package request

type CreateInventoryRequest struct {
	ProductID   uint32 `validate:"required" json:"product_id"`
	WarehouseID uint32 `validate:"required" json:"warehouse_id"`
	Quantity    int32  `validate:"required" json:"quantity"`
	Status      int8   `form:"default=1" json:"status"`
}

type UpdateInventoryRequest struct {
	ID          uint32 `validate:"required" json:"id"`
	ProductID   uint32 `validate:"required" json:"product_id"`
	WarehouseID uint32 `validate:"required" json:"warehouse_id"`
	Quantity    int32  `validate:"required" json:"quantity"`
	Status      int8   `json:"status"`
}
