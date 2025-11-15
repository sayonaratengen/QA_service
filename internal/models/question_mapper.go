package models

import "github.com/sayonaratengen/QA_service/internal/domain"

func ToGORMQuestion(q domain.Question) QuestionGORM {
	return QuestionGORM{
		ID:        q.ID,
		Text:      q.Text,
		CreatedAt: q.CreatedAt,
	}
}

func ToDomainQuestion(g QuestionGORM) domain.Question {
	answers := make([]domain.Answer, len(g.Answers))
	for i, a := range g.Answers {
		answers[i] = ToDomainAnswer(a)
	}

	return domain.Question{
		ID:        g.ID,
		Text:      g.Text,
		CreatedAt: g.CreatedAt,
	}
}
