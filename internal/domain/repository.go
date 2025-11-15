package domain

import "context"

type QuestionRepository interface {
	Create(ctx context.Context, q Question) (Question, error)
	GetByID(ctx context.Context, id int) (Question, error)
	GetAll(ctx context.Context) ([]Question, error)
	Delete(ctx context.Context, id int) error
}

type AnswerRepository interface {
	Create(ctx context.Context, a Answer) (Answer, error)
	GetByID(ctx context.Context, id int) (Answer, error)
	Delete(ctx context.Context, id int) error
	GetByQuestionID(ctx context.Context, questionID int) ([]Answer, error)
}
