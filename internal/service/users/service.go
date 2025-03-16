package users

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	// Get retuns user with omitted Password field (empty string)
	Get(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, req *model.CreateUserReq) error
	Exists(ctx context.Context, username string) (bool, error)
	UserWithEmailExists(ctx context.Context, email string) (bool, error)
	GetPasswordHash(ctx context.Context, username string) (string, error)
}

type service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) Service {
	return &service{
		db: db,
	}
}
