package models

import "github.com/sayonaratengen/QA_service/internal/domain"

func ToGORMAnswer(a domain.Answer) AnswerGORM {
	return AnswerGORM{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
	}
}

func ToDomainAnswer(a AnswerGORM) domain.Answer {
	return domain.Answer{
		ID:         a.ID,
		QuestionID: a.QuestionID,
		UserID:     a.UserID,
		Text:       a.Text,
		CreatedAt:  a.CreatedAt,
	}
}
