package quiz

import (
	"diana-quiz/internal/service/auth"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// authorize returns username, 200, nil on success,
// empty string, status code, error on failure
func (h handler) authorize(c *fiber.Ctx) (string, int, error) {
	accessToken := c.Cookies(accessTokenCookieName)
	refreshToken := c.Cookies(refreshTokenCookieName)
	if len(accessToken) == 0 || len(refreshToken) == 0 {
		return "", fiber.StatusUnauthorized, errors.New("user is not logged in")
	}

	username, refreshedUserTokens, err := h.authService.Authorize(c.Context(), accessToken, refreshToken)
	if err != nil {
		if errors.Is(err, auth.ErrNotAuthorized) {
			return "", fiber.StatusUnauthorized, err
		}
		return "", fiber.StatusInternalServerError, err
	}

	c.Cookie(&fiber.Cookie{
		Name:     accessTokenCookieName,
		Value:    refreshedUserTokens.AccessToken,
		HTTPOnly: true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshedUserTokens.RefreshToken,
		HTTPOnly: true,
	})

	return username, fiber.StatusOK, nil
}
