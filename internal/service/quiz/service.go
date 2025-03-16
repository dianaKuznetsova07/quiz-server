package quiz

import (
	"context"
	"diana-quiz/internal/db"
	"diana-quiz/internal/model"

	"github.com/danielblagy/go-utils/logger"
)

type Service interface {
	// Create creates new quiz and returns the created quiz's ID.
	Create(ctx context.Context, req *model.CreateQuizReq, username string) (int64, error)
	GetByID(ctx context.Context, quizID int64) (*model.GetQuizResponse, error)
	CompleteQuiz(ctx context.Context, quizID int64, username string, req *model.CompleteQuizReq) error
	GetQuizOwner(ctx context.Context, quizID int64) (string, error)
	GetQuizResults(ctx context.Context, quizID int64) (*model.GetQuizResultsResponse, error)
	GetUserQuizes(ctx context.Context, username string) (*model.GetUserQuizesResponse, error)
}

type service struct {
	logger       logger.Logger
	queryFactory db.QueryFactory
}

func NewService(
	logger logger.Logger,
	queryFactory db.QueryFactory,
) Service {
	return &service{
		logger:       logger,
		queryFactory: queryFactory,
	}
}
