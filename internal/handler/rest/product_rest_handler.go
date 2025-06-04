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

type productRestHandler struct {
	productUsecase usecase.ProductUsecase
	validator      *pkgValidator.XValidator
}

func InitProductRestHandler(app *fiber.App, productUsecase usecase.ProductUsecase, middleware middleware.Middleware, config *config.EnvConfig, validator *pkgValidator.XValidator) {
	handler := &productRestHandler{
		productUsecase: productUsecase,
		validator:      validator,
	}

	app.Post("/v1/product", middleware.ValidateToken(config.JwtPublicKey), handler.CreateProduct)
	app.Get("/v1/product/:id", middleware.ValidateToken(config.JwtPublicKey), handler.GetProductByID)
	app.Get("/v1/product", middleware.ValidateToken(config.JwtPublicKey), handler.GetAllProducts)
	app.Put("/v1/product/:id", middleware.ValidateToken(config.JwtPublicKey), handler.UpdateProduct)
	app.Delete("/v1/product/:id", middleware.ValidateToken(config.JwtPublicKey), handler.DeleteProduct)
}

func (h *productRestHandler) CreateProduct(c *fiber.Ctx) error {
	req := &dto.CreateProductReq{}
	if err := c.BodyParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	image, err := c.FormFile("image")
	if err != nil {
		return helper.RespError(c, err)
	}
	req.Image = image
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.productUsecase.CreateProduct(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "create product success")
}

func (h *productRestHandler) GetProductByID(c *fiber.Ctx) error {
	req := &dto.GetProductByIDReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.productUsecase.GetProductByID(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get product by id success")
}

func (h *productRestHandler) GetAllProducts(c *fiber.Ctx) error {
	res, err := h.productUsecase.GetAllProducts(c.UserContext())
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get all products success")
}

func (h *productRestHandler) UpdateProduct(c *fiber.Ctx) error {
	req := &dto.UpdateProductReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := c.BodyParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	image, err := c.FormFile("image")
	if err != nil {
		return helper.RespError(c, err)
	}
	req.Image = image
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.productUsecase.UpdateProduct(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "update product success")
}

func (h *productRestHandler) DeleteProduct(c *fiber.Ctx) error {
	req := &dto.DeleteProductReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	if err := h.productUsecase.DeleteProduct(c.UserContext(), *req); err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, nil, "delete product success")
}
