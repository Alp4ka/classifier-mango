package http

import (
	"github.com/Alp4ka/classifier-mango/pkg/security"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) mwContentChecker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		const APIKeyHeader = "x-api-key"

		providedKey := c.Get(APIKeyHeader)
		ok := security.CheckPasswordHash(providedKey, s.cfg.APIKey)
		if !ok {
			return c.SendStatus(fiber.StatusForbidden)
		}

		if c.Method() != fiber.MethodGet && c.Method() != fiber.MethodDelete && !c.Is("json") {
			return c.SendStatus(fiber.StatusUnsupportedMediaType)
		}

		return c.Next()
	}
}
