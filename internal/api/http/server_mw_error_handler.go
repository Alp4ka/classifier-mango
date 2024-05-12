package http

import "github.com/gofiber/fiber/v2"

type HandlerResp struct {
	Data    any    `json:"data"`
	Message string `json:"message,omitempty"`
}

func (s *Server) mwErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(HandlerResp{
		Message: err.Error(),
	})
}

var _ fiber.ErrorHandler = (*Server)(nil).mwErrorHandler
