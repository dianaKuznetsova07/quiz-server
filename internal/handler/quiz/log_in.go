package quiz

import (
	"diana-quiz/internal/model"
	"diana-quiz/internal/service/auth"
	"errors"

	"github.com/gofiber/fiber/v2"
)

const accessTokenCookieName = "diana-quiz-access-token"
const refreshTokenCookieName = "diana-quiz-refresh-token"

func (h *handler) logIn(c *fiber.Ctx) error {
	var login model.LoginReq
	if err := c.BodyParser(&login); err != nil {
		return err
	}

	if err := h.validate.Struct(login); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	userTokens, err := h.authService.LogIn(c.Context(), &login)
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		if errors.Is(err, auth.ErrIncorrectPassword) {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    userTokens.AccessToken,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    userTokens.RefreshToken,
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).SendString("successfully logged in")
}
