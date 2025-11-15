package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/sayonaratengen/QA_service/internal/middleware"
	"github.com/sayonaratengen/QA_service/pkg/logger"
)

func NewRouter(ctx context.Context, qh *QuestionHandler, ah *AnswerHandler, httpTimeout time.Duration) http.Handler {
	log := logger.FromContext(ctx)
	mux := http.NewServeMux()

	mux.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			qh.GetListAllQuestions(w, r)
		case http.MethodPost:
			qh.CreateQuestion(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, MsgMethodNotAllowed)
		}
	})

	mux.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDFromPath("/questions/", r)
		if err != nil {
			writeError(w, http.StatusBadRequest, MsgInvalidQuestionID)
			return
		}

		if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/answers") {
			ah.CreateAnswer(w, r, id)
			return
		}

		switch r.Method {
		case http.MethodGet:
			qh.GetQuestionByID(w, r, id)
		case http.MethodDelete:
			qh.DeleteQuestion(w, r, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, MsgMethodNotAllowed)
		}
	})
	
	mux.HandleFunc("/answers/", func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIDFromPath("/answers/", r)
		if err != nil {
			writeError(w, http.StatusBadRequest, MsgInvalidAnswerID)
			return
		}

		switch r.Method {
		case http.MethodGet:
			ah.GetAnswerByID(w, r, id)
		case http.MethodDelete:
			ah.DeleteAnswer(w, r, id)
		default:
			writeError(w, http.StatusMethodNotAllowed, MsgMethodNotAllowed)
		}
	})

	handlerWithLogger := middleware.RequestLogger(log, mux)
	return middleware.Timeout(httpTimeout, handlerWithLogger)
}
