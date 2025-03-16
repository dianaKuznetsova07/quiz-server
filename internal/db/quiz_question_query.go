package db

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/Masterminds/squirrel"
	"github.com/elgris/stom"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type QuizQuestionQuery interface {
	// Add returnes ID of created quiz record.
	Add(ctx context.Context, toAdd []*model.QuizQuestion) error
	GetByQuizID(ctx context.Context, quizID int64) ([]*model.QuizQuestion, error)
}

type quizQuestionQuery struct {
	runner pgxscan.Querier
}

var quizQuestionSelectColumns = stom.MustNewStom(model.QuizQuestion{}).SetTag(SelectTag).TagValues()

func newQuizQuestionQuery(runner pgxscan.Querier) QuizQuestionQuery {
	return &quizQuestionQuery{
		runner: runner,
	}
}

func (q *quizQuestionQuery) Add(ctx context.Context, toAdd []*model.QuizQuestion) error {
	qb := squirrel.
		Insert(model.GetQuizQuestionTableName()).
		Columns(
			model.QuizQuestionColumnQuizID,
			model.QuizQuestionColumnTitle,
			model.QuizQuestionColumnQuestionType,
			model.QuizQuestionColumnOptions,
			model.QuizQuestionColumnCreatedAt,
		).
		PlaceholderFormat(squirrel.Dollar)

	for _, e := range toAdd {
		qb = qb.Values(
			e.QuizID,
			e.Title,
			e.QuestionType,
			e.Options,
			e.CreatedAt,
		)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	rows, err := q.runner.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (q *quizQuestionQuery) GetByQuizID(ctx context.Context, quizID int64) ([]*model.QuizQuestion, error) {
	qb := squirrel.
		Select(quizQuestionSelectColumns...).
		From(model.GetQuizQuestionTableName()).
		Where(squirrel.Eq{
			model.QuizQuestionColumnQuizID: quizID,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dest []*model.QuizQuestion
	err = pgxscan.Select(ctx, q.runner, &dest, query, args...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
