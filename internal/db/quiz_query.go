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

type QuizQuery interface {
	// Add returnes ID of created quiz record.
	Add(ctx context.Context, title, username string) (int64, error)
	GetByID(ctx context.Context, quizID int64) (*model.Quiz, error)
	GetQuizOwner(ctx context.Context, quizID int64) (string, error)
	GetByOwner(ctx context.Context, username string) ([]*model.Quiz, error)
}

type quizQuery struct {
	runner pgxscan.Querier
}

var quizSelectColumns = stom.MustNewStom(model.Quiz{}).SetTag(SelectTag).TagValues()

func newQuizQuery(runner pgxscan.Querier) QuizQuery {
	return &quizQuery{
		runner: runner,
	}
}

func (q *quizQuery) Add(ctx context.Context, title, username string) (int64, error) {
	qb := squirrel.
		Insert(model.GetQuizTableName()).
		Columns(
			model.QuizColumnTitle,
			model.QuizColumnOwnerUsername,
			model.QuizColumnCreatedAt,
			model.QuizColumnUpdatedAt,
		).
		Values(
			title,
			username,
			time.Now(),
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", model.QuizColumnID)).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return 0, err
	}

	var createdQuizID int64
	if err := pgxscan.Get(ctx, q.runner, &createdQuizID, query, args...); err != nil {
		return 0, err
	}

	return createdQuizID, nil
}

func (q *quizQuery) GetByID(ctx context.Context, quizID int64) (*model.Quiz, error) {
	qb := squirrel.
		Select(quizSelectColumns...).
		From(model.GetQuizTableName()).
		Where(squirrel.Eq{
			model.QuizColumnID: quizID,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dest model.Quiz
	if err := pgxscan.Get(ctx, q.runner, &dest, query, args...); err != nil {
		return nil, err
	}

	return &dest, nil
}

func (q *quizQuery) GetQuizOwner(ctx context.Context, quizID int64) (string, error) {
	qb := squirrel.
		Select(model.QuizColumnOwnerUsername).
		From(model.GetQuizTableName()).
		Where(squirrel.Eq{
			model.QuizColumnID: quizID,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return "", err
	}

	var dest string
	if err := pgxscan.Get(ctx, q.runner, &dest, query, args...); err != nil {
		return "", err
	}

	return dest, nil
}

func (q *quizQuery) GetByOwner(ctx context.Context, username string) ([]*model.Quiz, error) {
	qb := squirrel.
		Select(quizSelectColumns...).
		From(model.GetQuizTableName()).
		Where(squirrel.Eq{
			model.QuizColumnOwnerUsername: username,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dest []*model.Quiz
	if err := pgxscan.Select(ctx, q.runner, &dest, query, args...); err != nil {
		return nil, err
	}

	return dest, nil
}
