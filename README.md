**Этот проект использует переменные окружения.  
Файл `.env` в репозиторий не коммитится намеренно и добавлен в `.gitignore`.**

_Скопировать `.env.example` в `.env` и при необходимости изменить параметры:_

```
cp .env.example .env
```

Поднять сервис и базу можно напрямую через docker-compose (но тогда нужно не забыть создать .env в корне проекта на основе .env.example):
```
docker compose up --build

либо

docker-compose up --build
```

_Далее в Taskfile для совместимости используется переменная DOCKER_COMPOSE, поэтому работает и с docker compose (v2), и с docker-compose (v1), но при запуске напрямую отталкивайтесь от вашей версии._

---

Миграции доступны вручную через:
```
docker compose exec app bash -c "GOOSE_DRIVER=postgres GOOSE_DBSTRING='host=db user=postgres password=pass dbname=qa sslmode=disable' GOOSE_MIGRATION_DIR=/app/migrations goose up <version>"
docker compose exec app bash -c "GOOSE_DRIVER=postgres GOOSE_DBSTRING='host=db user=postgres password=pass dbname=qa sslmode=disable' GOOSE_MIGRATION_DIR=/app/migrations goose down <version>"

либо

docker-compose exec app bash -c "GOOSE_DRIVER=postgres GOOSE_DBSTRING='host=db user=postgres password=pass dbname=qa sslmode=disable' GOOSE_MIGRATION_DIR=/app/migrations goose up <version>"
docker-compose exec app bash -c "GOOSE_DRIVER=postgres GOOSE_DBSTRING='host=db user=postgres password=pass dbname=qa sslmode=disable' GOOSE_MIGRATION_DIR=/app/migrations goose down <version>"
```
Все миграции для создания таблиц и индексов также применяются автоматически при запуске приложения.

---

**Также для удобства написан Taskfile.**

_Основные команды для Taskfile:_

1. Запуск тестов:
``task test
``
2. Собрать Docker-образ приложения:
``task build``
3. Собрать и поднять контейнеры (app + db):
``task up``
4. Остановить и удалить контейнеры:
``task down``
5. Применить ВСЕ миграции вручную через goose:
``task migrate``

_Для Taskfile нужен https://taskfile.dev/#/installation_

_При использовании task-команд .env создается автоматически (при его отсутствии)._


---
**Технические детали**

- **Go** 1.25, стандартная библиотека net/http для API
- **PostgreSQL** + GORM для работы с БД
- Миграции с помощью **goose**
- Логирование через **zap**, ошибки обрабатываются и возвращаются корректные HTTP-статусы
- Структура проекта модульная: отдельные пакеты для `handler`, `service`, `repository`, `domain`, `pkg`

**Особенности реализации ТЗ**
- Проверка существования вопроса перед созданием ответа
- Каскадное удаление реализовано через GORM `OnDelete:CASCADE`
- Константы для сообщений и ошибок вынесены в отдельные файлы для читаемости кода
- Используется Docker + docker-compose для локального поднятия проекта

---

**Для проекта написаны юнит-тесты, которые проверяют ключевую бизнес-логику сервиса:**

- невозможность создания ответа к несуществующему вопросу;
- корректную обработку нескольких ответов от одного и того же пользователя.

_Тесты не требуют поднятой базы данных: используются моки репозиториев._


