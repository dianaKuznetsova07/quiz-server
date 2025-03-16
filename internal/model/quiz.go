package model

import (
	"time"
)

const (
	QuizColumnID            = "id"
	QuizColumnTitle         = "title"
	QuizColumnOwnerUsername = "owner_username"
	QuizColumnCreatedAt     = "created_at"
	QuizColumnUpdatedAt     = "updated_at"
)

type Quiz struct {
	ID            int64     `db:"id" json:"id"`
	Title         string    `db:"title" json:"title"`
	OwnerUsername string    `db:"owner_username" json:"owner_username"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

func GetQuizTableName() string {
	return "quiz"
}
