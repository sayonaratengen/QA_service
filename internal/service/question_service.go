package service

import (
	"context"
	"fmt"

	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
)

type QuestionService struct {
	repo domain.QuestionRepository
}

func NewQuestionService(repo domain.QuestionRepository) QuestionServiceInterface {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) Create(ctx context.Context, q domain.Question) (domain.Question, error) {
	log := logger.FromContext(ctx)

	if q.Text == "" {
		log.Warn(msgQuestionValidation)
		return domain.Question{}, domain.ErrQuestionTextEmpty
	}

	createdQ, err := s.repo.Create(ctx, q)
	if err != nil {
		log.Error(msgQuestionCreateFail, zap.Error(err))
		return domain.Question{}, fmt.Errorf("%s: %w", msgQuestionCreateFail, err)
	}

	log.Info(msgQuestionCreated, zap.Int("id", createdQ.ID))
	return createdQ, nil
}

func (s *QuestionService) GetByID(ctx context.Context, id int) (domain.Question, error) {
	log := logger.FromContext(ctx)

	q, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Warn(msgQuestionNotFound, zap.Int("id", id))
		return domain.Question{}, fmt.Errorf("%s: %w", msgQuestionNotFound, err)
	}

	return q, nil
}

func (s *QuestionService) GetAll(ctx context.Context) ([]domain.Question, error) {
	log := logger.FromContext(ctx)

	questions, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error(msgQuestionGetAllFail, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgQuestionGetAllFail, err)
	}

	log.Info("список вопросов получен", zap.Int("count", len(questions)))
	return questions, nil
}

func (s *QuestionService) Delete(ctx context.Context, id int) error {
	log := logger.FromContext(ctx)

	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Error(msgQuestionDeleted, zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("%s: %w", msgQuestionDeleted, err)
	}

	log.Info(msgQuestionDeleted, zap.Int("id", id))
	return nil
}
