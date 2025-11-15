package mapper

import (
	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/internal/handler/dto"
)

func ToDomainQuestion(req dto.CreateQuestionRequest) domain.Question {
	return domain.Question{
		Text: req.Text,
	}
}

func ToQuestionResponse(q domain.Question) dto.QuestionResponse {
	return dto.QuestionResponse{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: toRFC3339ms(q.CreatedAt),
	}
}

func ToQuestionResponseList(questions []domain.Question) []dto.QuestionResponse {
	res := make([]dto.QuestionResponse, len(questions))
	for i, q := range questions {
		res[i] = ToQuestionResponse(q)
	}
	return res
}

func ToQuestionWithAnswersResponse(q domain.Question, answers []domain.Answer) dto.QuestionWithAnswersResponse {
	respAnswers := make([]dto.AnswerResponse, len(answers))
	for i, a := range answers {
		respAnswers[i] = ToAnswerResponse(a)
	}

	return dto.QuestionWithAnswersResponse{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: toRFC3339ms(q.CreatedAt),
		Answers:   respAnswers,
	}
}
