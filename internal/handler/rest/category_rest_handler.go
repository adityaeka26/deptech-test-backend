package rest

import (
	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/middleware"
	"github.com/adityaeka26/deptech-test-backend/internal/usecase"
	pkgError "github.com/adityaeka26/deptech-test-backend/pkg/error"
	"github.com/adityaeka26/deptech-test-backend/pkg/helper"
	pkgValidator "github.com/adityaeka26/deptech-test-backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type categoryRestHandler struct {
	categoryUsecase usecase.CategoryUsecase
	validator       *pkgValidator.XValidator
}

func InitCategoryRestHandler(app *fiber.App, categoryUsecase usecase.CategoryUsecase, middleware middleware.Middleware, config *config.EnvConfig, validator *pkgValidator.XValidator) {
	handler := &categoryRestHandler{
		categoryUsecase: categoryUsecase,
		validator:       validator,
	}

	app.Post("/v1/category", middleware.ValidateToken(config.JwtPublicKey), handler.CreateCategory)
	app.Get("/v1/category/:id", middleware.ValidateToken(config.JwtPublicKey), handler.GetCategoryByID)
	app.Get("/v1/category", middleware.ValidateToken(config.JwtPublicKey), handler.GetAllCategories)
	app.Put("/v1/category/:id", middleware.ValidateToken(config.JwtPublicKey), handler.UpdateCategory)
	app.Delete("/v1/category/:id", middleware.ValidateToken(config.JwtPublicKey), handler.DeleteCategory)
}

func (h *categoryRestHandler) CreateCategory(c *fiber.Ctx) error {
	req := &dto.CreateCategoryReq{}
	if err := c.BodyParser(&req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.categoryUsecase.CreateCategory(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "create category success")
}

func (h *categoryRestHandler) GetCategoryByID(c *fiber.Ctx) error {
	req := &dto.GetCategoryByIDReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.categoryUsecase.GetCategoryByID(c.UserContext(), req.ID)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get category by id success")
}

func (h *categoryRestHandler) GetAllCategories(c *fiber.Ctx) error {
	res, err := h.categoryUsecase.GetAllCategories(c.UserContext())
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get all categories success")
}

func (h *categoryRestHandler) UpdateCategory(c *fiber.Ctx) error {
	req := &dto.UpdateCategoryReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := c.BodyParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.categoryUsecase.UpdateCategory(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "update category success")
}

func (h *categoryRestHandler) DeleteCategory(c *fiber.Ctx) error {
	req := &dto.DeleteCategoryReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	if err := h.categoryUsecase.DeleteCategory(c.UserContext(), *req); err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, nil, "delete category success")
}
