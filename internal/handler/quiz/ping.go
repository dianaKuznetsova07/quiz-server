package quiz

import "github.com/gofiber/fiber/v2"

func (h *handler) ping(c *fiber.Ctx) error {
	return c.SendString("Hello from diana-quiz\n")
}
