package db

import (
	"context"

	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	msgDBConnectFail    = "ошибка подключения к БД"
	msgDBConnectSuccess = "успешное подключение к БД"
)

func NewPostgreSQL(ctx context.Context, dsn string) (*gorm.DB, error) {
	log := logger.FromContext(ctx)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(msgDBConnectFail, zap.Error(err))
		return nil, err
	}

	log.Info(msgDBConnectSuccess)
	return db, nil
}
