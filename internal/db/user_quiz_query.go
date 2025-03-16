package db

import (
	"context"
	"diana-quiz/internal/model"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/elgris/stom"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type UserQuizQuery interface {
	// Add returnes ID of created user quiz record.
	Add(ctx context.Context, quizID int64, username string) (int64, error)
	GetByQuizID(ctx context.Context, quizID int64) ([]*model.UserQuiz, error)
}

type userQuizQuery struct {
	runner pgxscan.Querier
}

var userQuizSelectColumns = stom.MustNewStom(model.UserQuiz{}).SetTag(SelectTag).TagValues()

func newUserQuizQuery(runner pgxscan.Querier) UserQuizQuery {
	return &userQuizQuery{
		runner: runner,
	}
}

func (q *userQuizQuery) Add(ctx context.Context, quizID int64, username string) (int64, error) {
	qb := squirrel.
		Insert(model.GetUserQuizTableName()).
		Columns(
			model.UserQuizColumnQuizID,
			model.UserQuizColumnUsername,
			model.UserQuizColumnFinishedAt,
		).
		Values(
			quizID,
			username,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", model.UserQuizColumnID)).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}

	var createdUserQuizID int64
	if err := pgxscan.Get(ctx, q.runner, &createdUserQuizID, query, args...); err != nil {
		return 0, err
	}

	return createdUserQuizID, nil
}

func (q *userQuizQuery) GetByQuizID(ctx context.Context, quizID int64) ([]*model.UserQuiz, error) {
	qb := squirrel.
		Select(userQuizSelectColumns...).
		From(model.GetUserQuizTableName()).
		Where(squirrel.Eq{
			model.UserQuizColumnQuizID: quizID,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dest []*model.UserQuiz
	if err := pgxscan.Select(ctx, q.runner, &dest, query, args...); err != nil {
		return nil, err
	}

	return dest, nil
}
