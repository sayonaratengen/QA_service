package main

import (
	"context"
	"log"

	"github.com/sayonaratengen/QA_service/internal/app"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
)

const (
	msgLoggerInitFail = "не удалось инициализировать логгер"
	msgConfigLoadFail = "не удалось загрузить конфиг"
	msgServerCreate   = "не удалось создать сервер"
	msgServerCrash    = "сервер завершил работу с ошибкой"
)

func main() {
	l, err := logger.InitLogger()
	if err != nil {
		log.Fatalf("%s: %v", msgLoggerInitFail, err)
	}
	defer l.Sync()

	ctx := context.Background()
	ctx = logger.WithContext(ctx, l)

	cfg, err := app.LoadConfig()
	if err != nil {
		l.Fatal(msgConfigLoadFail, zap.Error(err))
	}

	srv, err := app.NewServer(ctx, cfg)
	if err != nil {
		l.Fatal(msgServerCreate, zap.Error(err))
	}

	if err := srv.StartAndWait(ctx); err != nil {
		l.Fatal(msgServerCrash, zap.Error(err))
	}
}
