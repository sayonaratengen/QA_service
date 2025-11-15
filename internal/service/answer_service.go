package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
)

type AnswerService struct {
	repo         domain.AnswerRepository
	questionRepo domain.QuestionRepository
}

func NewAnswerService(answerRepo domain.AnswerRepository, questionRepo domain.QuestionRepository) AnswerServiceInterface {
	return &AnswerService{
		repo:         answerRepo,
		questionRepo: questionRepo,
	}
}

func (s *AnswerService) Create(ctx context.Context, a domain.Answer) (domain.Answer, error) {
	log := logger.FromContext(ctx)

	if a.QuestionID == 0 {
		log.Warn(msgAnswerValidationQ)
		return domain.Answer{}, domain.ErrAnswerQuestionIDZero
	}
	if a.UserID == "" {
		log.Warn(msgAnswerValidationU)
		return domain.Answer{}, domain.ErrAnswerUserIDEmpty
	}
	if a.Text == "" {
		log.Warn(msgAnswerValidationT)
		return domain.Answer{}, domain.ErrAnswerTextEmpty
	}

	_, err := s.questionRepo.GetByID(ctx, a.QuestionID)
	if err != nil {
		if errors.Is(err, domain.ErrQuestionNotFound) {
			log.Warn(msgAnswerQuestionNotFound, zap.Int("question_id", a.QuestionID))
			return domain.Answer{}, domain.ErrQuestionNotFound
		}
		return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerQuestionCheckFailed, err)
	}

	createdA, err := s.repo.Create(ctx, a)
	if err != nil {
		log.Error(msgAnswerCreateFail, zap.Error(err))
		return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerCreateFail, err)
	}

	log.Info(msgAnswerCreated, zap.Int("id", createdA.ID))
	return createdA, nil
}

func (s *AnswerService) GetByID(ctx context.Context, id int) (domain.Answer, error) {
	log := logger.FromContext(ctx)

	a, err := s.repo.GetByID(ctx, id)
	if err != nil {
		log.Warn(msgAnswerNotFound, zap.Int("id", id))
		return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerNotFound, err)
	}

	return a, nil
}

func (s *AnswerService) GetByQuestionID(ctx context.Context, questionID int) ([]domain.Answer, error) {
	log := logger.FromContext(ctx)

	answers, err := s.repo.GetByQuestionID(ctx, questionID)
	if err != nil {
		log.Error(msgAnswerGetByQFail, zap.Int("question_id", questionID), zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgAnswerGetByQFail, err)
	}

	log.Info(msgAnswerGetByQFail, zap.Int("question_id", questionID), zap.Int("count", len(answers)))
	return answers, nil
}

func (s *AnswerService) Delete(ctx context.Context, id int) error {
	log := logger.FromContext(ctx)

	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Error(msgAnswerDeleted, zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("%s: %w", msgAnswerDeleted, err)
	}

	log.Info(msgAnswerDeleted, zap.Int("id", id))
	return nil
}
