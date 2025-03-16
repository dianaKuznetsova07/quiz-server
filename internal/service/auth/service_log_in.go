package auth

import (
	"context"
	"diana-quiz/internal/model"
	"diana-quiz/internal/service/users"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrIncorrectPassword = errors.New("password is incorrect")

func (s *service) LogIn(ctx context.Context, req *model.LoginReq) (*model.UserTokens, error) {
	passwordHash, err := s.usersService.GetPasswordHash(ctx, req.Username)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrIncorrectPassword
		}
		return nil, errors.Wrap(err, "can't compare passwords")
	}

	accessToken, err := generateJwtToken(req.Username, accessTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate access token")
	}

	refreshToken, err := generateJwtToken(req.Username, refreshTokenDuration)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate refresh token")
	}

	return &model.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
