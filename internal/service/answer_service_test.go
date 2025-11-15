package service

import (
	"context"
	"testing"

	"github.com/sayonaratengen/QA_service/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestAnswerService_CreateAnswerForQuestion(t *testing.T) {
	ctx := context.Background()
	qRepo := NewQuestionRepositoryMock()
	aRepo := NewAnswerRepositoryMock()
	qSvc := NewQuestionService(qRepo)
	aSvc := NewAnswerService(aRepo, qRepo)

	q, _ := qSvc.Create(ctx, domain.Question{Text: "Вопрос для ответов"})

	a1, err := aSvc.Create(ctx, domain.Answer{QuestionID: q.ID, UserID: "user1", Text: "ответ1"})
	assert.NoError(t, err)
	a2, err := aSvc.Create(ctx, domain.Answer{QuestionID: q.ID, UserID: "user1", Text: "ответ2"})
	assert.NoError(t, err)

	assert.NotEqual(t, a1.ID, a2.ID)
}

func TestAnswerService_CreateAnswerForNonexistentQuestion(t *testing.T) {
	ctx := context.Background()
	qRepo := NewQuestionRepositoryMock()
	aRepo := NewAnswerRepositoryMock()
	aSvc := NewAnswerService(aRepo, qRepo)

	_, err := aSvc.Create(ctx, domain.Answer{QuestionID: 999, UserID: "user1", Text: "ответ"})
	assert.Error(t, err)
}
