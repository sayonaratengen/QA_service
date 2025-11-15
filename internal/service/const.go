package service

const (
	msgQuestionCreateFail = "не удалось создать вопрос"
	msgQuestionCreated    = "вопрос успешно создан"
	msgQuestionNotFound   = "вопрос не найден"
	msgQuestionGetAllFail = "не удалось получить список вопросов"
	msgQuestionDeleted    = "вопрос удалён"
	msgQuestionValidation = "создание вопроса: текст пустой"

	msgAnswerCreateFail          = "не удалось создать ответ"
	msgAnswerCreated             = "ответ успешно создан"
	msgAnswerNotFound            = "ответ не найден"
	msgAnswerGetByQFail          = "не удалось получить ответы по вопросу"
	msgAnswerDeleted             = "ответ удалён"
	msgAnswerValidationQ         = "создание ответа: QuestionID = 0"
	msgAnswerValidationU         = "создание ответа: UserID пустой"
	msgAnswerValidationT         = "создание ответа: текст пустой"
	msgAnswerQuestionNotFound    = "попытка создать ответ к несуществующему вопросу"
	msgAnswerQuestionCheckFailed = "ошибка проверки существования вопроса"
)
