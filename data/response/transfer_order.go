package response

import "time"

type TransferOrderResponse struct {
	ID                   uint32    `json:"id"`
	UserID               uint32    `json:"user_id"`
	ProductID            uint32    `json:"product_id"`
	SourceWarehouse      uint32    `json:"source_warehouse"`
	DestinationWarehouse uint32    `json:"destination_warehouse"`
	Status               string    `json:"status"`
	Quantity             int32     `json:"quantity"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	SuccessAt            time.Time `json:"success_at"`
	FailedAt             time.Time `json:"failed_at"`
}
