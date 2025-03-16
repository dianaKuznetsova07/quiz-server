package db

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/Masterminds/squirrel"
	"github.com/elgris/stom"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type UserQuizAnswerQuery interface {
	// Add returnes ID of created user quiz record.
	Add(ctx context.Context, answers []*model.UserQuizAnswer) error
	GetByUserQuizIDs(ctx context.Context, userQuizIDs []int64) (map[int64][]*model.UserQuizAnswer, error)
}

type userQuizAnswerQuery struct {
	runner pgxscan.Querier
}

var userQuizAnswerSelectColumns = stom.MustNewStom(model.UserQuizAnswer{}).SetTag(SelectTag).TagValues()

func newUserQuizAnswerQuery(runner pgxscan.Querier) UserQuizAnswerQuery {
	return &userQuizAnswerQuery{
		runner: runner,
	}
}

func (q *userQuizAnswerQuery) Add(ctx context.Context, answers []*model.UserQuizAnswer) error {
	qb := squirrel.
		Insert(model.GetUserQuizAnswerTableName()).
		Columns(
			model.UserQuizAnswerColumnUserQuizID,
			model.UserQuizAnswerColumnQuizQuestionID,
			model.UserQuizAnswerColumnOptionAnswer,
			model.UserQuizAnswerColumnTextAnswer,
		).
		PlaceholderFormat(squirrel.Dollar)

	for _, a := range answers {
		qb = qb.Values(
			a.UserQuizID,
			a.QuizQuestionID,
			a.OptionAnswer,
			a.TextAnswer,
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

func (q *userQuizAnswerQuery) GetByUserQuizIDs(ctx context.Context, userQuizIDs []int64) (map[int64][]*model.UserQuizAnswer, error) {
	qb := squirrel.
		Select(userQuizAnswerSelectColumns...).
		From(model.GetUserQuizAnswerTableName()).
		Where(squirrel.Eq{
			model.UserQuizAnswerColumnUserQuizID: userQuizIDs,
		}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dest []*model.UserQuizAnswer
	if err := pgxscan.Select(ctx, q.runner, &dest, query, args...); err != nil {
		return nil, err
	}

	userQuizAnswersMap := make(map[int64][]*model.UserQuizAnswer, len(userQuizIDs))
	for _, userQuizAnswer := range dest {
		userQuizAnswersMap[userQuizAnswer.UserQuizID] = append(userQuizAnswersMap[userQuizAnswer.UserQuizID], userQuizAnswer)
	}

	return userQuizAnswersMap, nil
}
