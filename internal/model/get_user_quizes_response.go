package model

import "time"

type GetUserQuizesResponse struct {
	Quizes []GetUserQuizesResponseQuiz `json:"quizes"`
}

type GetUserQuizesResponseQuiz struct {
	QuizID    int64     `json:"quiz_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
