package model

import "time"

type GetQuizResultsResponse struct {
	ID          int64                              `json:"id"`
	Title       string                             `json:"title"`
	UserResults []GetQuizResultsResponseUserResult `json:"user_results"`
}

type GetQuizResultsResponseUserResult struct {
	FinishedAt time.Time                                `json:"finished_at"`
	Username   string                                   `json:"username"`
	Answers    []GetQuizResultsResponseUserResultAnswer `json:"answers"`
}

type GetQuizResultsResponseUserResultAnswer struct {
	QuizQuestionID int64   `json:"quiz_question_id"`
	OptionAnswer   *string `json:"option_answer"`
	TextAnswer     *string `json:"text_answer"`
}
