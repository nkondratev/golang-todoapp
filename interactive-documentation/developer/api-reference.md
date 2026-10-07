# Справочник API

**Разделы:** [Главная](../index.md) | Пользователю: [Быстрый старт](../user/quick-start.md) · [Руководство](../user/user-guide.md) · [FAQ](../user/faq.md) | Разработчику: [Установка](../developer/installation.md) · [Архитектура](../developer/architecture.md) · [API](../developer/api-reference.md)

Справочник описывает HTTP-маршруты приложения. Формат данных — JSON, адрес по умолчанию — `http://localhost:8080`.

> **Важно:** маршруты и поля ниже приведены как типовые для TODO API. Актуальное описание смотрите в Swagger-документации (генерируется `swaggo`) и сверяйте с кодом слоя Handler.

## Маршруты

| Метод | Путь | Назначение |
|---|---|---|
| POST | `/users` | Создать пользователя |
| POST | `/tasks` | Создать задачу |
| GET | `/tasks` | Получить список задач |
| GET | `/tasks/{id}` | Получить задачу по идентификатору |
| PUT | `/tasks/{id}` | Изменить задачу или её статус |
| DELETE | `/tasks/{id}` | Удалить задачу |

## Модель задачи

| Поле | Тип | Описание |
|---|---|---|
| `id` | число | Идентификатор |
| `user_id` | число | Владелец задачи |
| `title` | строка | Название |
| `status` | строка | Статус (новая, в работе, выполнена) |
| `created_at` | дата-время | Дата создания |
| `closed_at` | дата-время | Дата закрытия, заполняется при завершении |

## Примеры

**Создать пользователя:**

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Иван"}'
```

**Создать задачу:**

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "title": "Подготовить отчёт"}'
```

**Закрыть задачу:**

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"status": "done"}'
```

**Удалить задачу:**

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

Ожидаемый результат: успешный код ответа (2xx) и корректная запись операции в логах (`docker compose logs`). Эти запросы используются для [проверки после установки](installation.md#проверка-установки). Как ответственность распределена по слоям — в [архитектуре](architecture.md).
