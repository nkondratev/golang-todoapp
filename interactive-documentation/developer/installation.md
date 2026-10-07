# Установка и настройка

**Разделы:** [Главная](../index.md) | Пользователю: [Быстрый старт](../user/quick-start.md) · [Руководство](../user/user-guide.md) · [FAQ](../user/faq.md) | Разработчику: [Установка](../developer/installation.md) · [Архитектура](../developer/architecture.md) · [API](../developer/api-reference.md)

Раздел отвечает на вопросы: что нужно для запуска, как установить приложение и как его настроить. Для понимания устройства системы см. [архитектуру](architecture.md).

## Системные требования

| Параметр | Требование |
|---|---|
| Операционная система | Linux |
| Процессор | Intel Core i3 и выше |
| Оперативная память | 6 ГБ |
| Свободное место | 10 ГБ |
| Сеть | доступ к серверу по локальной сети |
| Программы | Docker, Docker Compose, Git |

Для разработки без Docker дополнительно: Go 1.27.1 и PostgreSQL 18. Используемые библиотеки: `zap` (логирование), `swaggo` (Swagger), `pgx` (драйвер PostgreSQL), `envconfig` (конфигурация из переменных окружения).

## Установка

1. Проверьте наличие Docker и Docker Compose:

   ```bash
   docker --version
   docker compose version
   ```

2. Получите исходные файлы проекта (с `Dockerfile` и `docker-compose.yml`):

   ```bash
   git clone https://github.com/nkondratev/golang-todoapp.git
   cd golang-todoapp
   ```

3. Соберите и запустите контейнеры:

   ```bash
   make env-up
   make migrate-up
   make env-port-forward
   make todoapp-run
   ```

4. Проверьте, что контейнеры работают:

   ```bash
   make ps
   ```

## Настройка

Параметры подключения к БД, порт API и переменные окружения задаются в `docker-compose.yml` и файле конфигурации. Читаются библиотекой `envconfig`.

| Переменная | Назначение | Пример |
|---|---|---|
| `APP_PORT` | Порт API | `8080` |
| `DB_HOST` | Адрес PostgreSQL | `db` |
| `DB_PORT` | Порт PostgreSQL | `5432` |
| `DB_USER` | Пользователь БД | `todo` |
| `DB_PASSWORD` | Пароль БД | задаётся администратором |
| `DB_NAME` | Имя базы данных | `todo` |

> **Важно:** точные имена переменных сверьте с вашим `docker-compose.yml` и структурой конфигурации в коде.

## Проверка установки

```bash
curl http://localhost:8080/tasks
```

Ожидается ответ с кодом 200 и списком задач (возможно, пустым). Затем создайте, получите, измените и удалите тестовую задачу — примеры в [справочнике API](api-reference.md).

