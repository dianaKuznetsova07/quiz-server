package quiz

import (
	"context"
	"diana-quiz/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var ErrQuizNotFound = errors.New("quiz not found")

func (s *service) GetByID(ctx context.Context, quizID int64) (*model.GetQuizResponse, error) {
	quiz, err := s.queryFactory.NewQuizQuery(s.queryFactory.GetConnectionPool()).GetByID(ctx, quizID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrQuizNotFound
		}
		return nil, errors.Wrap(err, "can't get quiz")
	}

	response := &model.GetQuizResponse{
		ID:        quiz.ID,
		Owner:     quiz.OwnerUsername,
		Title:     quiz.Title,
		CreatedAt: quiz.CreatedAt,
	}

	quizQuestions, err := s.queryFactory.NewQuizQuestionQuery(s.queryFactory.GetConnectionPool()).GetByQuizID(ctx, quizID)
	if err != nil {
		return nil, errors.Wrap(err, "can't get quiz questions")
	}
	response.Questions = make([]model.GetQuizResponseQuestion, 0, len(quizQuestions))

	for _, quizQuestion := range quizQuestions {
		response.Questions = append(response.Questions, model.GetQuizResponseQuestion{
			ID:           quizQuestion.ID,
			Title:        quizQuestion.Title,
			QuestionType: quizQuestion.QuestionType,
			Options:      quizQuestion.Options,
		})
	}

	return response, nil
}
