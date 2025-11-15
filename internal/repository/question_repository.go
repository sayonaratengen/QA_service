package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/internal/models"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type QuestionRepositoryGORM struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) domain.QuestionRepository {
	return &QuestionRepositoryGORM{db: db}
}

func (r *QuestionRepositoryGORM) Create(ctx context.Context, q domain.Question) (domain.Question, error) {
	log := logger.FromContext(ctx)

	gormQ := models.ToGORMQuestion(q)
	if err := r.db.Create(&gormQ).Error; err != nil {
		log.Error(msgQuestionCreateFail, zap.Error(err))
		return domain.Question{}, fmt.Errorf("%s: %w", msgQuestionCreateFail, err)
	}

	log.Info(msgQuestionCreated, zap.Int("id", gormQ.ID))
	return models.ToDomainQuestion(gormQ), nil
}

func (r *QuestionRepositoryGORM) GetByID(ctx context.Context, id int) (domain.Question, error) {
	log := logger.FromContext(ctx)

	var gormQ models.QuestionGORM
	if err := r.db.Preload("Answers").First(&gormQ, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Question{}, fmt.Errorf("%s: %w", msgQuestionGetFail, domain.ErrQuestionNotFound)
		}
		log.Error(msgQuestionGetFail, zap.Int("id", id), zap.Error(err))
		return domain.Question{}, fmt.Errorf("%s: %w", msgQuestionGetFail, err)
	}

	return models.ToDomainQuestion(gormQ), nil
}

func (r *QuestionRepositoryGORM) GetAll(ctx context.Context) ([]domain.Question, error) {
	log := logger.FromContext(ctx)

	var gormQs []models.QuestionGORM
	if err := r.db.Find(&gormQs).Error; err != nil {
		log.Error(msgQuestionsGetAllFail, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgQuestionsGetAllFail, err)
	}

	questions := make([]domain.Question, 0, len(gormQs))
	for _, q := range gormQs {
		questions = append(questions, models.ToDomainQuestion(q))
	}

	return questions, nil
}

func (r *QuestionRepositoryGORM) Delete(ctx context.Context, id int) error {
	log := logger.FromContext(ctx)

	if err := r.db.Delete(&models.QuestionGORM{}, id).Error; err != nil {
		log.Error(msgQuestionDeleteFail, zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("%s: %w", msgQuestionDeleteFail, err)
	}

	log.Info(msgQuestionDeleted, zap.Int("id", id))
	return nil
}
