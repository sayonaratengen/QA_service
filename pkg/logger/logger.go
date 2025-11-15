package logger

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type ctxKeyType string

const (
	loggerKey       ctxKeyType = "logger"
	msgFallbackLog             = "не удалось создать fallback логгер"
	msgEmptyContext            = "контекст пуст, создан fallback логгер"
	msgNoLogger                = "логгер не найден в контексте, создан fallback logger"
)

func InitLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func WithContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		l, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("%s: %v", msgFallbackLog, err)
		}
		l.Info(msgEmptyContext)
		return l
	}

	l, ok := ctx.Value(loggerKey).(*zap.Logger)
	if !ok || l == nil {
		newLogger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("%s: %v", msgFallbackLog, err)
		}
		newLogger.Info(msgNoLogger)
		return newLogger
	}
	return l
}
