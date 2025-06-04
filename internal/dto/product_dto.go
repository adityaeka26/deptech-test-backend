package dto

import "mime/multipart"

type CreateProductReq struct {
	Name        string                `form:"name" validate:"required"`
	Description string                `form:"description" validate:"required"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
	CategoryID  uint                  `form:"category_id" validate:"required"`
	Stock       uint                  `form:"stock" validate:"required"`
}

type CreateProductRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	CategoryID  uint   `json:"category_id"`
	Stock       uint   `json:"stock"`
}

type GetProductByIDReq struct {
	ID uint `param:"id" validate:"required"`
}

type GetProductByIDRes struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	ImageUrl    string             `json:"image_url"`
	Stock       uint               `json:"stock"`
	Category    GetCategoryByIDRes `json:"category"`
}

type UpdateProductReq struct {
	ID          uint                  `param:"id" validate:"required"`
	Name        string                `form:"name" validate:"required"`
	Description string                `form:"description" validate:"required"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
	CategoryID  uint                  `form:"category_id" validate:"required"`
	Stock       uint                  `form:"stock" validate:"required"`
}

type UpdateProductRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	CategoryID  uint   `json:"category_id"`
	Stock       uint   `json:"stock"`
}

type DeleteProductReq struct {
	ID uint `param:"id" validate:"required"`
}
