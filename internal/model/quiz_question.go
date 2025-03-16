package model

import "time"

const (
	QuizQuestionColumnID           = "id"
	QuizQuestionColumnQuizID       = "quiz_id"
	QuizQuestionColumnTitle        = "title"
	QuizQuestionColumnQuestionType = "question_type"
	QuizQuestionColumnOptions      = "options"
	QuizQuestionColumnCreatedAt    = "created_at"
)

type QuizQuestion struct {
	ID           int64            `db:"id" json:"id"`
	QuizID       int64            `db:"quiz_id" json:"quiz_id"`
	Title        string           `db:"title" json:"title"`
	QuestionType QuizQuestionType `db:"question_type" json:"question_type"`
	Options      []string         `db:"options" json:"options"`
	CreatedAt    time.Time        `db:"created_at" json:"created_at"`
}

type QuizQuestionType string

// QuizQuestionType values
const (
	QuizQuestionTypeText   QuizQuestionType = "text"
	QuizQuestionTypeChoice QuizQuestionType = "choice"
)

var ValidQuestionTypeMap = map[QuizQuestionType]struct{}{
	QuizQuestionTypeText:   {},
	QuizQuestionTypeChoice: {},
}

func GetQuizQuestionTableName() string {
	return "quiz_question"
}
