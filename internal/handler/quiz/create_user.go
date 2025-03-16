package quiz

import (
	"diana-quiz/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) createUser(c *fiber.Ctx) error {
	var req model.CreateUserReq
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validateCreateUserReq(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	exists, err := h.usersService.Exists(c.Context(), req.Username)
	if err != nil {
		return err
	}
	if exists {
		return c.Status(fiber.StatusConflict).SendString("username has already been taken")
	}

	userWithEmailExists, err := h.usersService.UserWithEmailExists(c.Context(), req.Email)
	if err != nil {
		return err
	}
	if userWithEmailExists {
		return c.Status(fiber.StatusConflict).SendString("email is already in use by some user")
	}

	if err := h.usersService.Create(c.Context(), &req); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) validateCreateUserReq(req *model.CreateUserReq) error {
	if err := h.validate.Struct(req); err != nil {
		return err
	}

	return nil
}
