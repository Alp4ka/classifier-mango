package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type hFinishParams struct {
	SessionID uuid.UUID `form:"sessionID"`
}

type hFinishResp struct{}

func (s *Server) hFinish(c *fiber.Ctx) error {
	ctx := c.UserContext()

	params := new(hFinishParams)
	if err := c.QueryParser(params); err != nil {
		return err
	}

	err := s.cfg.CoreClient.ReleaseSession(ctx, params.SessionID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

var _ fiber.Handler = (*Server)(nil).hStart
