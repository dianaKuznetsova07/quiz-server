package model

import "time"

const (
	UserQuizColumnID         = "id"
	UserQuizColumnQuizID     = "quiz_id"
	UserQuizColumnUsername   = "username"
	UserQuizColumnFinishedAt = "finished_at"
)

type UserQuiz struct {
	ID         int64     `db:"id" json:"id"`
	QuizID     int64     `db:"quiz_id" json:"quiz_id"`
	Username   string    `db:"username" json:"username"`
	FinishedAt time.Time `db:"finished_at" json:"finished_at"`
}

func GetUserQuizTableName() string {
	return "user_quiz"
}
