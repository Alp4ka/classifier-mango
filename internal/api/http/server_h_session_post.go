package http

import (
	classifiermango "github.com/Alp4ka/classifier-mango"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type hStartReq struct {
	Agent string `json:"agent"`
}
type hStartResp struct {
	SessionID uuid.UUID `json:"sessionID"`
}

func (s *Server) hStart(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hStartReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	sessionID, err := s.cfg.CoreClient.AcquireSession(ctx, req.Agent, classifiermango.AppName)
	if err != nil {
		return err
	}

	return c.
		Status(fiber.StatusOK).
		JSON(
			HandlerResp{
				Data: &hStartResp{SessionID: sessionID},
			},
		)
}

var _ fiber.Handler = (*Server)(nil).hStart
