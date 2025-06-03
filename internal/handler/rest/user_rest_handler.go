package rest

import (
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/usecase"
	pkgError "github.com/adityaeka26/deptech-test-backend/pkg/error"
	"github.com/adityaeka26/deptech-test-backend/pkg/helper"
	pkgValidator "github.com/adityaeka26/deptech-test-backend/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type userRestHandler struct {
	userUsecase usecase.UserUsecase
	validator   *pkgValidator.XValidator
}

func InitUserRestHandler(app *fiber.App, userUsecase usecase.UserUsecase, validator *pkgValidator.XValidator) {
	handler := &userRestHandler{
		userUsecase: userUsecase,
		validator:   validator,
	}

	app.Post("/v1/user", handler.CreateUser)
	app.Get("/v1/user/:id", handler.GetUserByID)
}

func (h *userRestHandler) CreateUser(c *fiber.Ctx) error {
	req := &dto.CreateUserReq{}
	if err := c.BodyParser(&req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.userUsecase.CreateUser(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "create user success")
}

func (h *userRestHandler) GetUserByID(c *fiber.Ctx) error {
	req := &dto.GetUserByIDReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.userUsecase.GetUserByID(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get user by id success")
}
