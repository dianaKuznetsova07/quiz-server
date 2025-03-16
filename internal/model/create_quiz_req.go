package model

type CreateQuizReq struct {
	Title     string `json:"title" validate:"required,min=2,max=512"`
	Questions []CreateQuizReqQuestion
}

type CreateQuizReqQuestion struct {
	Title string           `json:"title" validate:"required,min=2,max=512"`
	Type  QuizQuestionType `json:"type"`
	// Options will be empty for 'text' type questions.
	Options []string `json:"options"`
}
