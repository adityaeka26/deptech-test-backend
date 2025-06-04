package dto

type CreateCategoryReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateCategoryRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetCategoryByIDReq struct {
	ID uint `param:"id" validate:"required"`
}

type GetCategoryByIDRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryReq struct {
	ID          uint   `param:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateCategoryRes struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteCategoryReq struct {
	ID uint `param:"id" validate:"required"`
}
