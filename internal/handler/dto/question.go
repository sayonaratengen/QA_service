package dto

type CreateQuestionRequest struct {
	Text string `json:"text"`
}

type QuestionResponse struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type QuestionWithAnswersResponse struct {
	ID        int              `json:"id"`
	Text      string           `json:"text"`
	CreatedAt string           `json:"created_at"`
	Answers   []AnswerResponse `json:"answers"`
}
