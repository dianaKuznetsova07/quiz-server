package users

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) GetPasswordHash(ctx context.Context, username string) (string, error) {
	var passwordHash string
	err := pgxscan.Get(ctx, s.db, &passwordHash, fmt.Sprintf("select password_hash from users where username = '%s'", username))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrUserNotFound
		}
		return "", errors.Wrap(err, "can't get user's password_hash from db")
	}

	return passwordHash, nil
}
