package model

import "time"

type GetQuizResponse struct {
	ID        int64                     `json:"id"`
	Owner     string                    `json:"owner"`
	IsOwner   bool                      `json:"is_owner"`
	Title     string                    `json:"title"`
	CreatedAt time.Time                 `json:"created_at"`
	Questions []GetQuizResponseQuestion `json:"questions"`
}

type GetQuizResponseQuestion struct {
	ID           int64            `json:"id"`
	Title        string           `json:"title"`
	QuestionType QuizQuestionType `json:"question_type"`
	Options      []string         `json:"options"`
}
