package middleware

import (
	"strings"

	pkgError "github.com/adityaeka26/deptech-test-backend/pkg/error"
	"github.com/adityaeka26/deptech-test-backend/pkg/helper"
	"github.com/adityaeka26/deptech-test-backend/pkg/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type middleware struct {
	redis *redis.Redis
}

type Middleware interface {
	ValidateToken(jwtPublicKey string) fiber.Handler
}

func NewMiddleware(redis *redis.Redis) *middleware {
	return &middleware{
		redis: redis,
	}
}

func (m *middleware) ValidateToken(jwtPublicKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtPublicKey = strings.ReplaceAll(jwtPublicKey, "\\n", "\n")

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(jwtPublicKey))
		if err != nil {
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if len(token) <= 0 {
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		red, err := m.redis.RedisClient.Get(c.UserContext(), token).Result()
		if err != nil && err.Error() != "redis: nil" {
			return err
		}
		if red != "" {
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}

		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})
		if err != nil {
			return helper.RespError(c, pkgError.UnauthorizedError("unauthorized"))
		}
		claims := parseToken.Claims.(jwt.MapClaims)

		c.Locals("id", claims["data"].(map[string]any)["id"])
		c.Locals("email", claims["data"].(map[string]any)["email"])

		return c.Next()
	}
}
