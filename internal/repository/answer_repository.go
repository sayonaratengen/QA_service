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

type AnswerRepositoryGORM struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) domain.AnswerRepository {
	return &AnswerRepositoryGORM{db: db}
}

func (r *AnswerRepositoryGORM) Create(ctx context.Context, a domain.Answer) (domain.Answer, error) {
	log := logger.FromContext(ctx)

	gormA := models.ToGORMAnswer(a)
	if err := r.db.Create(&gormA).Error; err != nil {
		log.Error(msgAnswerCreateFail, zap.Error(err))
		return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerCreateFail, err)
	}

	log.Info(msgAnswerCreated, zap.Int("id", gormA.ID))
	return models.ToDomainAnswer(gormA), nil
}

func (r *AnswerRepositoryGORM) GetByID(ctx context.Context, id int) (domain.Answer, error) {
	log := logger.FromContext(ctx)

	var gormA models.AnswerGORM
	if err := r.db.First(&gormA, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerGetFail, domain.ErrAnswerNotFound)
		}
		log.Error(msgAnswerGetFail, zap.Int("id", id), zap.Error(err))
		return domain.Answer{}, fmt.Errorf("%s: %w", msgAnswerGetFail, err)
	}

	return models.ToDomainAnswer(gormA), nil
}

func (r *AnswerRepositoryGORM) Delete(ctx context.Context, id int) error {
	log := logger.FromContext(ctx)

	if err := r.db.Delete(&models.AnswerGORM{}, id).Error; err != nil {
		log.Error(msgAnswerDeleteFail, zap.Int("id", id), zap.Error(err))
		return fmt.Errorf("%s: %w", msgAnswerDeleteFail, err)
	}

	log.Info(msgAnswerDeleted, zap.Int("id", id))
	return nil
}

func (r *AnswerRepositoryGORM) GetByQuestionID(ctx context.Context, questionID int) ([]domain.Answer, error) {
	log := logger.FromContext(ctx)

	var gormAs []models.AnswerGORM
	if err := r.db.Where("question_id = ?", questionID).Find(&gormAs).Error; err != nil {
		log.Error(msgAnswersByQuestionFail, zap.Int("question_id", questionID), zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgAnswersByQuestionFail, err)
	}

	answers := make([]domain.Answer, 0, len(gormAs))
	for _, a := range gormAs {
		answers = append(answers, models.ToDomainAnswer(a))
	}

	return answers, nil
}
