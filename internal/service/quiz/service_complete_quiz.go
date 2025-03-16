package quiz

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (s *service) CompleteQuiz(ctx context.Context, quizID int64, username string, req *model.CompleteQuizReq) error {
	// get quiz
	quiz, err := s.GetByID(ctx, quizID)
	if err != nil {
		return errors.Wrap(err, "can't get quiz")
	}

	if validateErr := s.validateQuizAnswers(req, quiz); validateErr != nil {
		return validateErr
	}

	txRunner, err := s.queryFactory.GetConnectionPool().Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "can't begin tx")
	}

	err = s.saveUserAnswer(ctx, txRunner, quizID, username, req)
	if err != nil {
		if rollbackErr := txRunner.Rollback(ctx); rollbackErr != nil {
			s.logger.ErrorKV("can't rollback tx", "err", rollbackErr.Error())
		}
		return err
	}

	if commitErr := txRunner.Commit(ctx); commitErr != nil {
		return errors.Wrap(commitErr, "can't commit tx")
	}

	return nil
}

func (s *service) validateQuizAnswers(req *model.CompleteQuizReq, quiz *model.GetQuizResponse) error {
	if len(req.Answers) != len(quiz.Questions) {
		return errors.New("len answers != len questions")
	}

	questionsMap := make(map[int64]model.GetQuizResponseQuestion, len(quiz.Questions))
	for _, q := range quiz.Questions {
		questionsMap[q.ID] = q
	}

	for _, a := range req.Answers {
		if len(a.TextAnswer) == 0 && len(a.OptionAnswer) == 0 {
			errors.New("no answer provided")
		}

		q, ok := questionsMap[a.QuestionID]
		if !ok {
			return errors.New("answer to non-existent question")
		}

		if len(a.TextAnswer) > 0 && q.QuestionType == model.QuizQuestionTypeChoice {
			return errors.New("wrong type of answer to question")
		}

		if len(a.OptionAnswer) > 0 && q.QuestionType == model.QuizQuestionTypeText {
			return errors.New("wrong type of answer to question")
		}
	}

	return nil
}

func (s *service) saveUserAnswer(ctx context.Context, txRunner pgxscan.Querier, quizID int64, username string, req *model.CompleteQuizReq) error {
	createdUserQuizID, err := s.queryFactory.NewUserQuizQuery(txRunner).Add(ctx, quizID, username)
	if err != nil {
		return errors.Wrap(err, "can't add user quiz")
	}

	toAdd := make([]*model.UserQuizAnswer, 0, len(req.Answers))
	for _, a := range req.Answers {
		var optionAnswer *string
		if len(a.OptionAnswer) > 0 {
			optionAnswer = &a.OptionAnswer
		}
		var textAnswer *string
		if len(a.TextAnswer) > 0 {
			textAnswer = &a.TextAnswer
		}

		toAdd = append(toAdd, &model.UserQuizAnswer{
			UserQuizID:     createdUserQuizID,
			QuizQuestionID: a.QuestionID,
			OptionAnswer:   optionAnswer,
			TextAnswer:     textAnswer,
		})
	}

	err = s.queryFactory.NewUserQuizAnswerQuery(txRunner).Add(ctx, toAdd)
	if err != nil {
		return errors.Wrap(err, "can't add user quiz answers")
	}

	return nil
}
