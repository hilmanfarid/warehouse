package request

type CreateProductRequest struct {
	Name   string  `validate:"required,min=1,max=100" json:"name"`
	Code   string  `validate:"required,min=1,max=100" json:"code"`
	Price  float64 `validate:"required" json:"price"`
	Status int8    `form:"default=1" json:"status"`
}

type UpdateProductRequest struct {
	ID     uint32  `validate:"required"`
	Name   string  `validate:"required,min=1,max=100" json:"name"`
	Code   string  `validate:"required,min=1,max=100" json:"code"`
	Price  float64 `validate:"required" json:"price"`
	Status int8    `json:"status"`
}
