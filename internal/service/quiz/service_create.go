package quiz

import (
	"context"
	"diana-quiz/internal/model"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s *service) Create(ctx context.Context, req *model.CreateQuizReq, username string) (int64, error) {
	txRunner, err := s.queryFactory.GetConnectionPool().Begin(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "can't begin tx")
	}

	quizID, err := s.create(ctx, txRunner, req, username)
	if err != nil {
		if rollbackErr := txRunner.Rollback(ctx); rollbackErr != nil {
			s.logger.ErrorKV("can't rollback tx", "err", rollbackErr.Error())
		}
		return 0, err
	}

	if commitErr := txRunner.Commit(ctx); commitErr != nil {
		return 0, errors.Wrap(commitErr, "can't commit tx")
	}

	return quizID, nil
}

func (s *service) create(ctx context.Context, txRunner pgxscan.Querier, req *model.CreateQuizReq, username string) (int64, error) {
	// add quiz record

	quizID, err := s.queryFactory.NewQuizQuery(txRunner).Add(ctx, req.Title, username)
	if err != nil {
		return 0, errors.Wrap(err, "can't add quiz")
	}

	// add quiz questions

	quizQuestionsToAdd := make([]*model.QuizQuestion, 0, len(req.Questions))

	for _, questionReq := range req.Questions {
		quizQuestionToAdd := &model.QuizQuestion{
			QuizID:       quizID,
			Title:        questionReq.Title,
			QuestionType: questionReq.Type,
			CreatedAt:    time.Now(),
		}

		if questionReq.Type == model.QuizQuestionTypeChoice {
			quizQuestionToAdd.Options = questionReq.Options
		}

		quizQuestionsToAdd = append(quizQuestionsToAdd, quizQuestionToAdd)
	}

	if addErr := s.queryFactory.NewQuizQuestionQuery(txRunner).Add(ctx, quizQuestionsToAdd); addErr != nil {
		return 0, errors.Wrap(addErr, "can't add quiz questions")
	}

	return quizID, nil
}
