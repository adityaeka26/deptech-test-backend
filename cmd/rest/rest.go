package rest

import (
	"fmt"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/adityaeka26/deptech-test-backend/internal/handler/rest"
	"github.com/adityaeka26/deptech-test-backend/internal/middleware"
	"github.com/adityaeka26/deptech-test-backend/internal/usecase"
	pkgValidator "github.com/adityaeka26/deptech-test-backend/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ServeRest(config *config.EnvConfig, userUsecase usecase.UserUsecase, middleware middleware.Middleware) error {
	app := fiber.New()

	rest.InitUserRestHandler(app, userUsecase, middleware, config, &pkgValidator.XValidator{
		Validator: &validator.Validate{},
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))

	return nil
}
