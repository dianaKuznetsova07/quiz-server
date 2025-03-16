package auth

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/pkg/errors"
)

func (s *service) refreshTokens(ctx context.Context, refreshToken string) (string, *model.UserTokens, error) {
	username, err := parseJwtToken(refreshToken)
	if err != nil {
		if errors.Is(err, errTokenExpired) {
			return "", nil, errors.Wrap(ErrNotAuthorized, "refresh token has expired")
		}
		if errors.Is(err, errInvalidToken) {
			return "", nil, errors.Wrap(ErrNotAuthorized, "refresh token is invalid")
		}
		return "", nil, err
	}

	newAccessToken, err := generateJwtToken(username, accessTokenDuration)
	if err != nil {
		return "", nil, errors.Wrap(err, "can't generate access token")
	}

	newRefreshToken, err := generateJwtToken(username, refreshTokenDuration)
	if err != nil {
		return "", nil, errors.Wrap(err, "can't generate refresh token")
	}

	return username, &model.UserTokens{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
