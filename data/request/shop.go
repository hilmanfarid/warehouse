package request

type CreateShopRequest struct {
	Name   string `validate:"required,min=1,max=100" json:"name"`
	Status int8   `form:"default=1" json:"status"`
}

type UpdateShopRequest struct {
	ID     uint32 `validate:"required"`
	Name   string `validate:"required,min=1,max=100" json:"name"`
	Status int8   `json:"status"`
}
