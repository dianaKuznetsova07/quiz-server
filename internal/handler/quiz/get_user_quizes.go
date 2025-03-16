package quiz

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) getUserQuizes(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	response, err := h.quizService.GetUserQuizes(c.Context(), username)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
