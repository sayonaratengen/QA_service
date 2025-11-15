package dto

type CreateAnswerRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

type AnswerResponse struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	UserID     string `json:"user_id"`
	Text       string `json:"text"`
	CreatedAt  string `json:"created_at"`
}
