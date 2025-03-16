package quiz

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/pkg/errors"
)

func (s *service) GetUserQuizes(ctx context.Context, username string) (*model.GetUserQuizesResponse, error) {
	quizes, err := s.queryFactory.NewQuizQuery(s.queryFactory.GetConnectionPool()).GetByOwner(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "can't get quizes by username")
	}

	response := &model.GetUserQuizesResponse{
		Quizes: make([]model.GetUserQuizesResponseQuiz, 0, len(quizes)),
	}

	for _, quiz := range quizes {
		response.Quizes = append(response.Quizes, model.GetUserQuizesResponseQuiz{
			QuizID:    quiz.ID,
			Title:     quiz.Title,
			CreatedAt: quiz.CreatedAt,
		})
	}

	return response, nil
}
