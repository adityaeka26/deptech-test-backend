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

type transactionRestHandler struct {
	transactionUsecase usecase.TransactionUsecase
	validator          *pkgValidator.XValidator
}

func InitTransactionRestHandler(app *fiber.App, transactionUsecase usecase.TransactionUsecase, middleware middleware.Middleware, config *config.EnvConfig, validator *pkgValidator.XValidator) {
	handler := &transactionRestHandler{
		transactionUsecase: transactionUsecase,
		validator:          validator,
	}

	app.Post("/v1/transaction", middleware.ValidateToken(config.JwtPublicKey), handler.CreateTransaction)
	app.Get("/v1/transaction", middleware.ValidateToken(config.JwtPublicKey), handler.GetAllTransactions)
}

func (h *transactionRestHandler) CreateTransaction(c *fiber.Ctx) error {
	req := &dto.CreateTransactionReq{}
	if err := c.BodyParser(&req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	req.UserID = uint(c.Locals("id").(float64))
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.transactionUsecase.CreateTransaction(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "create transaction success")
}

func (h *transactionRestHandler) GetAllTransactions(c *fiber.Ctx) error {
	res, err := h.transactionUsecase.GetAllTransactions(c.UserContext())
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get all transactions success")
}
