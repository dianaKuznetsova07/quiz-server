package quiz

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (s *service) GetQuizOwner(ctx context.Context, quizID int64) (string, error) {
	quizOwner, err := s.queryFactory.NewQuizQuery(s.queryFactory.GetConnectionPool()).GetQuizOwner(ctx, quizID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", ErrQuizNotFound
		}
		return "", errors.Wrap(err, "can't get quiz owner")
	}

	return quizOwner, nil
}
