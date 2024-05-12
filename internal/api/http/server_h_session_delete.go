package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *Server) hFinish(c *fiber.Ctx) error {
	ctx := c.UserContext()

	paramSessionID := c.Params("session_id")
	parsedSessionID, err := uuid.Parse(paramSessionID)
	if err != nil {
		return err
	}

	err = s.cfg.CoreClient.ReleaseSession(ctx, parsedSessionID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

var _ fiber.Handler = (*Server)(nil).hStart
