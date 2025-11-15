package mapper

import (
	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/internal/handler/dto"
)

func ToDomainAnswer(req dto.CreateAnswerRequest, questionID int) domain.Answer {
	return domain.Answer{
		QuestionID: questionID,
		UserID:     req.UserID,
		Text:       req.Text,
	}
}

func ToAnswerResponse(a domain.Answer) dto.AnswerResponse {
	return dto.AnswerResponse{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  toRFC3339ms(a.CreatedAt),
	}
}
