package dto

import "time"

type CreateTransactionReq struct {
	Type   string                      `json:"type" validate:"required,oneof=in out"`
	UserID uint                        `validate:"required"`
	Items  []CreateTransactionReqItems `json:"items" validate:"required,dive"`
}

type CreateTransactionReqItems struct {
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

type GetAllTransactionsRes struct {
	ID        uint                         `json:"id"`
	User      GetAllTransactionsResUser    `json:"user"`
	Type      string                       `json:"type"`
	CreatedAt time.Time                    `json:"created_at"`
	Items     []GetAllTransactionsResItems `json:"items"`
}

type GetAllTransactionsResUser struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type GetAllTransactionsResItems struct {
	ID       uint                              `json:"id"`
	Product  GetAllTransactionsResItemsProduct `json:"product"`
	Quantity uint                              `json:"quantity"`
}

type GetAllTransactionsResItemsProduct struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	CategoryID  uint   `json:"category_id"`
	Stock       uint   `json:"stock"`
}
