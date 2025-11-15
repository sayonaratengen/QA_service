package service

import (
	"context"

	"github.com/sayonaratengen/QA_service/internal/domain"
)

type QuestionServiceInterface interface {
	Create(ctx context.Context, q domain.Question) (domain.Question, error)
	GetByID(ctx context.Context, id int) (domain.Question, error)
	GetAll(ctx context.Context) ([]domain.Question, error)
	Delete(ctx context.Context, id int) error
}

type AnswerServiceInterface interface {
	Create(ctx context.Context, a domain.Answer) (domain.Answer, error)
	GetByID(ctx context.Context, id int) (domain.Answer, error)
	Delete(ctx context.Context, id int) error
	GetByQuestionID(ctx context.Context, questionID int) ([]domain.Answer, error)
}
