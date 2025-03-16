package model

const (
	UserQuizAnswerColumnID             = "id"
	UserQuizAnswerColumnUserQuizID     = "user_quiz_id"
	UserQuizAnswerColumnQuizQuestionID = "quiz_question_id"
	UserQuizAnswerColumnOptionAnswer   = "option_answer"
	UserQuizAnswerColumnTextAnswer     = "text_answer"
)

type UserQuizAnswer struct {
	ID             int64   `db:"id" json:"id"`
	UserQuizID     int64   `db:"user_quiz_id" json:"user_quiz_id"`
	QuizQuestionID int64   `db:"quiz_question_id" json:"quiz_question_id"`
	OptionAnswer   *string `db:"option_answer" json:"option_answer"`
	TextAnswer     *string `db:"text_answer" json:"text_answer"`
}

func GetUserQuizAnswerTableName() string {
	return "user_quiz_answer"
}
