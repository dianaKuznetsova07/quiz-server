package users

import (
	"context"
	"diana-quiz/internal/model"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s service) UserWithEmailExists(ctx context.Context, email string) (bool, error) {
	var user model.User
	err := pgxscan.Get(ctx, s.db, &user, fmt.Sprintf("select username, email, full_name from users where email = '%s'", email))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, errors.Wrap(err, "can't check if user with email exists")
	}

	return true, nil
}
