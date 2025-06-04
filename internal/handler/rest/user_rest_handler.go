package rest

import (
	"strings"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/dto"
	"github.com/adityaeka26/deptech-test-backend/internal/middleware"
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

func InitUserRestHandler(app *fiber.App, userUsecase usecase.UserUsecase, middleware middleware.Middleware, config *config.EnvConfig, validator *pkgValidator.XValidator) {
	handler := &userRestHandler{
		userUsecase: userUsecase,
		validator:   validator,
	}

	app.Post("/v1/user", handler.CreateUser)
	app.Get("/v1/user/:id", handler.GetUserByID)
	app.Put("/v1/user/:id", middleware.ValidateToken(config.JwtPublicKey), handler.UpdateUser)
	app.Delete("/v1/user/:id", middleware.ValidateToken(config.JwtPublicKey), handler.DeleteUser)
	app.Get("/v1/user", handler.GetAllUsers)
	app.Post("/v1/user/login", handler.LoginUser)
	app.Post("/v1/user/logout", middleware.ValidateToken(config.JwtPublicKey), handler.LogoutUser)
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

func (h *userRestHandler) UpdateUser(c *fiber.Ctx) error {
	req := &dto.UpdateUserReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := c.BodyParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.userUsecase.UpdateUser(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "update user success")
}

func (h *userRestHandler) DeleteUser(c *fiber.Ctx) error {
	req := &dto.DeleteUserReq{}
	if err := c.ParamsParser(req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	if err := h.userUsecase.DeleteUser(c.UserContext(), *req); err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, nil, "delete user success")
}

func (h *userRestHandler) GetAllUsers(c *fiber.Ctx) error {
	res, err := h.userUsecase.GetAllUsers(c.UserContext())
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "get all users success")
}

func (h *userRestHandler) LoginUser(c *fiber.Ctx) error {
	req := &dto.LoginUserReq{}
	if err := c.BodyParser(&req); err != nil {
		return helper.RespError(c, pkgError.BadRequest(err.Error()))
	}
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	res, err := h.userUsecase.LoginUser(c.UserContext(), *req)
	if err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, res, "login user success")
}

func (h *userRestHandler) LogoutUser(c *fiber.Ctx) error {
	req := &dto.LogoutUserReq{}
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	if len(token) <= 0 {
		return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
	}
	req.Token = token
	if err := h.validator.Validate(req); err != nil {
		return helper.RespError(c, err)
	}

	if err := h.userUsecase.LogoutUser(c.UserContext(), *req); err != nil {
		return helper.RespError(c, err)
	}
	return helper.RespSuccess(c, nil, "logout user success")
}
