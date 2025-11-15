package app

const (
	msgEnvLoadFail       = "не удалось загрузить .env файл"
	msgEnvParseFail      = "не удалось распарсить переменные окружения"
	msgUnsupportedDriver = "неподдерживаемый драйвер БД"

	msgDSNBuildFail      = "не удалось сформировать DSN"
	msgDBConnectFail     = "не удалось подключиться к БД"
	msgDBConnected       = "БД успешно подключена"
	msgDBGetSQL          = "не удалось получить sql.DB из GORM"
	msgDBPingFail        = "БД недоступна"
	msgMigrationsFail    = "не удалось применить миграции"
	msgMigrationsDone    = "Миграции успешно выполнены"
	msgServerStart       = "сервер завершил работу с ошибкой"
	msgServerRunning     = "Сервер стартует"
	msgServerStopSig     = "Получен сигнал"
	msgServerShutdown    = "Сервер остановлен корректно"
	msgDBCloseFail       = "Ошибка при закрытии БД"
	msgServerShutdownErr = "Ошибка при завершении работы сервера"
)
