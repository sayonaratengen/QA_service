package service

import (
	"context"
	"testing"

	"github.com/sayonaratengen/QA_service/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestQuestionService_CreateAndGet(t *testing.T) {
	ctx := context.Background()
	qRepo := NewQuestionRepositoryMock()
	svc := NewQuestionService(qRepo)

	q, err := svc.Create(ctx, domain.Question{Text: "Вопрос 1"})
	assert.NoError(t, err)
	assert.NotZero(t, q.ID)

	got, err := svc.GetByID(ctx, q.ID)
	assert.NoError(t, err)
	assert.Equal(t, q.Text, got.Text)
}
