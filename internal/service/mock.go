package service

import (
	"context"
	"errors"
	"sync"

	"github.com/sayonaratengen/QA_service/internal/domain"
)

type QuestionRepositoryMock struct {
	data   map[int]domain.Question
	lastID int
	mu     sync.Mutex
}

func NewQuestionRepositoryMock() *QuestionRepositoryMock {
	return &QuestionRepositoryMock{
		data: make(map[int]domain.Question),
	}
}

func (r *QuestionRepositoryMock) Create(ctx context.Context, q domain.Question) (domain.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lastID++
	q.ID = r.lastID
	r.data[q.ID] = q
	return q, nil
}

func (r *QuestionRepositoryMock) GetByID(ctx context.Context, id int) (domain.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	q, ok := r.data[id]
	if !ok {
		return domain.Question{}, errors.New("вопрос не найден")
	}
	return q, nil
}

func (r *QuestionRepositoryMock) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return errors.New("вопрос не найден")
	}
	delete(r.data, id)
	return nil
}

func (r *QuestionRepositoryMock) GetAll(ctx context.Context) ([]domain.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := make([]domain.Question, 0, len(r.data))
	for _, q := range r.data {
		list = append(list, q)
	}
	return list, nil
}

type AnswerRepositoryMock struct {
	data   map[int]domain.Answer
	lastID int
	mu     sync.Mutex
}

func NewAnswerRepositoryMock() *AnswerRepositoryMock {
	return &AnswerRepositoryMock{
		data: make(map[int]domain.Answer),
	}
}

func (r *AnswerRepositoryMock) Create(ctx context.Context, a domain.Answer) (domain.Answer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lastID++
	a.ID = r.lastID
	r.data[a.ID] = a
	return a, nil
}

func (r *AnswerRepositoryMock) GetByID(ctx context.Context, id int) (domain.Answer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.data[id]
	if !ok {
		return domain.Answer{}, errors.New("ответ не найден")
	}
	return a, nil
}

func (r *AnswerRepositoryMock) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return errors.New("ответ не найден")
	}
	delete(r.data, id)
	return nil
}

func (r *AnswerRepositoryMock) GetByQuestionID(ctx context.Context, questionID int) ([]domain.Answer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := []domain.Answer{}
	for _, a := range r.data {
		if a.QuestionID == questionID {
			list = append(list, a)
		}
	}
	return list, nil
}
