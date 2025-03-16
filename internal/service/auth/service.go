package auth

import (
	"context"
	"diana-quiz/internal/model"
	"diana-quiz/internal/service/users"
)

type Service interface {
	LogIn(ctx context.Context, login *model.LoginReq) (*model.UserTokens, error)
	// Authorize returns username if successfully authenticated & a refreshed pair of tokens.
	Authorize(ctx context.Context, accessToken, refreshToken string) (string, *model.UserTokens, error)
}

type service struct {
	usersService users.Service
}

func NewService(usersService users.Service) Service {
	return &service{
		usersService: usersService,
	}
}
