package quiz

import (
	"diana-quiz/internal/model"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) createQuiz(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	var req model.CreateQuizReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validateCreateQuizReq(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	quizID, err := h.quizService.Create(c.Context(), &req, username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(&model.CreateQuizResponse{QuizID: quizID})
}

func (h *handler) validateCreateQuizReq(req *model.CreateQuizReq) error {
	if err := h.validate.Struct(req); err != nil {
		return err
	}

	if len(req.Questions) == 0 {
		return errors.New("questions field must not be empty")
	}

	for _, question := range req.Questions {
		if _, ok := model.ValidQuestionTypeMap[question.Type]; !ok {
			return errors.New("invalid question type")
		}
	}

	return nil
}
