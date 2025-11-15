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

type QuestionHandler struct {
	questions service.QuestionServiceInterface
	answers   service.AnswerServiceInterface
}

func NewQuestionHandler(qs service.QuestionServiceInterface, as service.AnswerServiceInterface) *QuestionHandler {
	return &QuestionHandler{
		questions: qs,
		answers:   as,
	}
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log := logger.FromContext(r.Context())
			log.Warn(MsgBodyCloseFailed, zap.Error(err))
		}
	}()

	var req dto.CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, MsgInvalidRequest)
		return
	}

	q, err := h.questions.Create(r.Context(), mapper.ToDomainQuestion(req))
	if err != nil {
		if errors.Is(err, domain.ErrQuestionTextEmpty) {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}

	writeJSON(w, http.StatusCreated, mapper.ToQuestionResponse(q))
}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request, id int) {
	q, err := h.questions.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrQuestionNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}

	answers, err := h.answers.GetByQuestionID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}

	writeJSON(w, http.StatusOK, mapper.ToQuestionWithAnswersResponse(q, answers))
}

func (h *QuestionHandler) GetListAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.questions.GetAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}

	writeJSON(w, http.StatusOK, mapper.ToQuestionResponseList(questions))
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.questions.Delete(r.Context(), id); err != nil {
		if errors.Is(err, domain.ErrQuestionNotFound) {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, MsgInternalError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
