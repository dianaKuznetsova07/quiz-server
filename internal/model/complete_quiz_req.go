package model

type CompleteQuizReq struct {
	Answers []CompleteQuizReqAnswer `json:"answers"`
}

type CompleteQuizReqAnswer struct {
	QuestionID   int64  `json:"question_id"`
	TextAnswer   string `json:"text_answer"`
	OptionAnswer string `json:"option_answer"`
}
