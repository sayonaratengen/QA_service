package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"
)

const (
	msgRequestTimeout = "время обработки запроса превышено"
)

func Timeout(duration time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), duration)
		defer cancel()

		done := make(chan struct{})
		go func() {
			next.ServeHTTP(w, r.WithContext(ctx))
			close(done)
		}()

		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				http.Error(w, msgRequestTimeout, http.StatusGatewayTimeout)
			}
		case <-done:
		}
	})
}
