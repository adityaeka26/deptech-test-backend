package dto

import "time"

type CreateTransactionReq struct {
	Type   string                     `json:"type" validate:"required,oneof=in out"`
	UserID uint                       `validate:"required"`
	Items  []CreateTransactionReqItem `json:"items" validate:"required,dive"`
}

type CreateTransactionReqItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
}

type CreateTransactionRes struct {
	ID        uint                       `json:"id"`
	UserID    uint                       `json:"user_id"`
	Type      string                     `json:"type"`
	CreatedAt time.Time                  `json:"created_at"`
	Items     []CreateTransactionResItem `json:"items"`
}

type CreateTransactionResItem struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
