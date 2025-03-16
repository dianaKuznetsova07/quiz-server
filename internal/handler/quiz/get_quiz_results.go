package quiz

import (
	"diana-quiz/internal/service/quiz"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) getQuizResults(c *fiber.Ctx) error {
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

	quizOwnerUsername, err := h.quizService.GetQuizOwner(c.Context(), int64(quizID))
	if err != nil {
		if errors.Is(err, quiz.ErrQuizNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if quizOwnerUsername != username {
		return c.Status(fiber.StatusUnauthorized).SendString("user is not the quiz owner")
	}

	response, err := h.quizService.GetQuizResults(c.Context(), int64(quizID))
	if err != nil {
		return err
	}

	return c.JSON(response)
}
