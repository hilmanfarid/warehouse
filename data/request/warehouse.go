package request

type CreateWarehouseRequest struct {
	ShopID uint32 `validate:"required" json:"shop_id"`
	Name   string `validate:"required,min=1,max=100" json:"name"`
	Status int8   `form:"default=1" json:"status"`
}

type UpdateWarehouseRequest struct {
	ID     uint32 `validate:"required"`
	ShopID uint32 `validate:"required" json:"shop_id"`
	Name   string `validate:"required,min=1,max=100" json:"name"`
	Status int8   `json:"status"`
}
