package domain

import "errors"

var (
	ErrQuestionTextEmpty = errors.New("ошибка валидации: текст вопроса не может быть пустым")
	ErrQuestionNotFound  = errors.New("вопрос не найден")

	ErrAnswerQuestionIDZero = errors.New("ошибка валидации: question_id не может быть 0")
	ErrAnswerUserIDEmpty    = errors.New("ошибка валидации: user_id не может быть пустым")
	ErrAnswerTextEmpty      = errors.New("ошибка валидации: текст ответа не может быть пустым")
	ErrAnswerNotFound       = errors.New("ответ не найден")
)
