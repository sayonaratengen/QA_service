package middleware

import (
	"net/http"

	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
)

const (
	msgIncomingRequest = "incoming request"
)

func RequestLogger(log *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.WithContext(r.Context(), log)
		r = r.WithContext(ctx)

		log.Info(msgIncomingRequest,
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)

		next.ServeHTTP(w, r)
	})
}
