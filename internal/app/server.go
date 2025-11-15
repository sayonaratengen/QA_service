package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sayonaratengen/QA_service/internal/handler"
	"github.com/sayonaratengen/QA_service/internal/repository"
	"github.com/sayonaratengen/QA_service/internal/service"
	"github.com/sayonaratengen/QA_service/pkg/db"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
	dbClose    func() error
	timeout    time.Duration
}

func NewServer(ctx context.Context, cfg *Config) (*Server, error) {
	log := logger.FromContext(ctx)

	dsn, err := cfg.DSN()
	if err != nil {
		log.Error(msgDSNBuildFail, zap.Error(err))
		return nil, err
	}

	database, err := db.NewPostgreSQL(ctx, dsn)
	if err != nil {
		log.Error(msgDBConnectFail, zap.Error(err))
		return nil, err
	}
	log.Info(msgDBConnected)

	sqlDB, err := database.DB()
	if err != nil {
		log.Error(msgDBGetSQL, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgDBGetSQL, err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		log.Error(msgDBPingFail, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgDBPingFail, err)
	}

	if err := goose.Up(sqlDB, cfg.MigrationsPath); err != nil {
		log.Error(msgMigrationsFail, zap.Error(err))
		return nil, fmt.Errorf("%s: %w", msgMigrationsFail, err)
	}
	log.Info(msgMigrationsDone)

	qr := repository.NewQuestionRepository(database)
	ar := repository.NewAnswerRepository(database)

	qs := service.NewQuestionService(qr)
	as := service.NewAnswerService(ar, qr)

	qh := handler.NewQuestionHandler(qs, as)
	ah := handler.NewAnswerHandler(as)

	router := handler.NewRouter(ctx, qh, ah, cfg.HTTPTimeout)

	httpSrv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:      router,
		ReadTimeout:  cfg.HTTPTimeout,
		WriteTimeout: cfg.HTTPTimeout,
	}

	return &Server{
		httpServer: httpSrv,
		dbClose:    sqlDB.Close,
		timeout:    cfg.ShutdownTimeout,
	}, nil
}

func (s *Server) StartAndWait(ctx context.Context) error {
	log := logger.FromContext(ctx)

	defer func() {
		if err := s.dbClose(); err != nil {
			log.Error(msgDBCloseFail, zap.Error(err))
		}
	}()

	log.Info(msgServerRunning, zap.String("адрес", s.httpServer.Addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error, 1)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("%s: %w", msgServerStart, err)
		}
	}()

	select {
	case sig := <-quit:
		log.Info(msgServerStopSig, zap.String("signal", sig.String()))
	case err := <-errCh:
		return err
	}

	shutdownCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error(msgServerShutdownErr, zap.Error(err))
		return err
	}

	log.Info(msgServerShutdown)
	return nil
}
