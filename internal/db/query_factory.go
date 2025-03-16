package db

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QueryFactory interface {
	GetConnectionPool() *pgxpool.Pool
	NewQuizQuery(runner pgxscan.Querier) QuizQuery
	NewQuizQuestionQuery(runner pgxscan.Querier) QuizQuestionQuery
	NewUserQuizQuery(runner pgxscan.Querier) UserQuizQuery
	NewUserQuizAnswerQuery(runner pgxscan.Querier) UserQuizAnswerQuery
}

type queryFactory struct {
	connectionPool *pgxpool.Pool
}

func NewQueryFactory(connectionPool *pgxpool.Pool) QueryFactory {
	return &queryFactory{
		connectionPool: connectionPool,
	}
}

func (f *queryFactory) GetConnectionPool() *pgxpool.Pool {
	return f.connectionPool
}

func (f *queryFactory) NewQuizQuery(runner pgxscan.Querier) QuizQuery {
	return newQuizQuery(runner)
}

func (f *queryFactory) NewQuizQuestionQuery(runner pgxscan.Querier) QuizQuestionQuery {
	return newQuizQuestionQuery(runner)
}

func (f *queryFactory) NewUserQuizQuery(runner pgxscan.Querier) UserQuizQuery {
	return newUserQuizQuery(runner)
}

func (f *queryFactory) NewUserQuizAnswerQuery(runner pgxscan.Querier) UserQuizAnswerQuery {
	return newUserQuizAnswerQuery(runner)
}
