package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sayonaratengen/QA_service/internal/domain"
	"github.com/sayonaratengen/QA_service/internal/handler/dto"
	"github.com/sayonaratengen/QA_service/internal/handler/mapper"
	"github.com/sayonaratengen/QA_service/internal/service"
	"github.com/sayonaratengen/QA_service/pkg/logger"

	"go.uber.org/zap"
)

type AnswerHandler struct {
	service service.AnswerServiceInterface
}

func NewAnswerHandler(service service.AnswerServiceInterface) *AnswerHandler {
	return &AnswerHandler{service: service}
}

func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request, questionID int) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log := logger.FromContext(r.Context())
			log.Warn(MsgBodyCloseFailed, zap.Error(err))
		}
	}()

	var req dto.CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, MsgInvalidRequest)
		return
	}

	a, err := h.service.Create(r.Context(), mapper.ToDomainAnswer(req, questionID))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrAnswerTextEmpty),
			errors.Is(err, domain.ErrAnswerUserIDEmpty),
			errors.Is(err, domain.ErrAnswerQuestionIDZero):
			writeError(w, http.StatusBadRequest, err.Error())
			return
		case errors.Is(err, domain.ErrQuestionNotFound):
			writeError(w, http.StatusNotFound, err.Error())
			return
		default:
			writeError(w, http.StatusInternalServerError, MsgInternalError)
			return
		}
	}

	writeJSON(w, http.StatusCreated, mapper.ToAnswerResponse(a))
}

func (h *AnswerHandler) GetAnswerByID(w http.ResponseWriter, r *http.Request, id int) {
	a, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrAnswerNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}

	writeJSON(w, http.StatusOK, mapper.ToAnswerResponse(a))
}

func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, domain.ErrAnswerNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
