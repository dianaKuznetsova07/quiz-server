package quiz

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) logOut(c *fiber.Ctx) error {
	username, statusCode, err := h.authorize(c)
	if err != nil {
		return c.Status(statusCode).SendString(err.Error())
	}

	// delete cookies
	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-time.Second),
	})
	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-time.Second),
	})

	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("%s logged out", username))
}
