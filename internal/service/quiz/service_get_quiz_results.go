package quiz

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/pkg/errors"
)

func (s *service) GetQuizResults(ctx context.Context, quizID int64) (*model.GetQuizResultsResponse, error) {
	quiz, err := s.GetByID(ctx, quizID)
	if err != nil {
		return nil, err
	}

	response := &model.GetQuizResultsResponse{
		ID:    quizID,
		Title: quiz.Title,
	}

	userQuizes, err := s.queryFactory.NewUserQuizQuery(s.queryFactory.GetConnectionPool()).GetByQuizID(ctx, quizID)
	if err != nil {
		return nil, errors.Wrap(err, "can't get user quizes")
	}

	userQuizIDs := make([]int64, 0, len(userQuizes))
	for _, uq := range userQuizes {
		userQuizIDs = append(userQuizIDs, uq.ID)
	}

	userAnswersMap, err := s.queryFactory.NewUserQuizAnswerQuery(s.queryFactory.GetConnectionPool()).GetByUserQuizIDs(ctx, userQuizIDs)
	if err != nil {
		return nil, errors.Wrap(err, "can't get user quizes answers")
	}

	response.UserResults = make([]model.GetQuizResultsResponseUserResult, 0, len(userQuizes))
	for _, uq := range userQuizes {
		userResult := model.GetQuizResultsResponseUserResult{
			Username:   uq.Username,
			FinishedAt: uq.FinishedAt,
		}

		for _, ua := range userAnswersMap[uq.ID] {
			userResult.Answers = append(userResult.Answers, model.GetQuizResultsResponseUserResultAnswer{
				QuizQuestionID: ua.QuizQuestionID,
				OptionAnswer:   ua.OptionAnswer,
				TextAnswer:     ua.TextAnswer,
			})
		}

		response.UserResults = append(response.UserResults, userResult)
	}

	return response, nil
}
