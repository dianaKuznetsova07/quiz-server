package quiz

import (
	"diana-quiz/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) completeQuiz(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	quizID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}

	if quizID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("id is not valid")
	}

	var req model.CompleteQuizReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err = h.quizService.CompleteQuiz(c.Context(), int64(quizID), username, &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).SendString("quiz completed")
}
