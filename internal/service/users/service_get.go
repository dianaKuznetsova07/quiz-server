package users

import (
	"context"
	"diana-quiz/internal/model"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrUserNotFound = errors.New("user not found")

func (s service) Get(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := pgxscan.Get(ctx, s.db, &user, fmt.Sprintf("select username, email, full_name from users where username = '%s'", username))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "can't get user from db")
	}

	return &user, nil
}
