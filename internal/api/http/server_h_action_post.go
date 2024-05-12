package http

import (
	"errors"

	"github.com/Alp4ka/classifier-mango/internal/interactions/core"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type hActionPostReq struct {
	SessionID uuid.UUID `json:"sessionID"`
	UserInput string    `json:"userInput"`
}

type hActionPostResp struct {
	Action   core.Action `json:"action"`
	Response string      `json:"response"`
}

func (s *Server) hActionPost(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(hActionPostReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	handler, err := s.coreManager.Process(ctx, req.SessionID)
	if err != nil {
		return err
	}

	err = handler.GetInputHandler().Handle(&core.ProcessInput{
		UserInput: req.UserInput,
		RequestID: uuid.New(),
	})
	if err != nil {
		return errors.Join(err, handler.Error())
	}

	output, err := handler.GetOutputHandler().Await()
	if err != nil {
		return errors.Join(err, handler.Error())
	}

	return c.Status(fiber.StatusOK).JSON(
		HandlerResp{
			Data: &hActionPostResp{
				Action:   output.Action,
				Response: output.UserResponse,
			},
		},
	)
}

var _ fiber.Handler = (*Server)(nil).hActionPost
