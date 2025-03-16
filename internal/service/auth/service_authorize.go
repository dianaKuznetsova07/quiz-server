package auth

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/pkg/errors"
)

var ErrNotAuthorized = errors.New("not authorized")

func (s *service) Authorize(ctx context.Context, accessToken, refreshToken string) (string, *model.UserTokens, error) {
	username, refreshedUserTokens, err := s.getUsernameWithRefreshFallback(ctx, accessToken, refreshToken)
	if err != nil {
		return "", nil, err
	}

	exists, err := s.usersService.Exists(ctx, username)
	if err != nil {
		return "", nil, err
	}
	if !exists {
		return "", nil, errors.Wrap(ErrNotAuthorized, "user doesn't exist")
	}

	return username, refreshedUserTokens, nil
}

func (s *service) getUsernameWithRefreshFallback(ctx context.Context, accessToken, refreshToken string) (string, *model.UserTokens, error) {
	username, err := parseJwtToken(accessToken)
	if err == nil {
		return username, &model.UserTokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
	}

	if !errors.Is(err, errTokenExpired) && !errors.Is(err, errInvalidToken) {
		return "", nil, err
	}

	return s.refreshTokens(ctx, refreshToken)
}
