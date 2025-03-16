package quiz

import (
	"diana-quiz/internal/service/auth"
	"diana-quiz/internal/service/quiz"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *handler) getQuiz(c *fiber.Ctx) error {
	username, err := h.getUsernameIfExists(c)
	if err != nil {
		return err
	}

	quizID, err := c.ParamsInt("id", -1)
	if err != nil {
		return err
	}

	if quizID <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("id is not valid")
	}

	response, err := h.quizService.GetByID(c.Context(), int64(quizID))
	if err != nil {
		if errors.Is(err, quiz.ErrQuizNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	response.IsOwner = username == response.Owner

	return c.JSON(response)
}

func (h *handler) getUsernameIfExists(c *fiber.Ctx) (string, error) {
	username, _, err := h.authorize(c)
	if err != nil {
		if errors.Is(err, auth.ErrNotAuthorized) {
			return "", nil
		}

		return "", errors.Wrap(err, "can't get user")
	}

	return username, nil
}
