# Документация «Задачи онлайн»

Интерактивная документация веб-приложения **«Задачи онлайн»** (TODO API) в формате Markdown.
Подготовлена для двух ролей: **пользователь** и **разработчик/администратор**.

Начните с файла [index.md](index.md).

## Структура проекта

```
interactive-documentation/
├── README.md
├── index.md
├── mkdocs.yml
├── user/
│   ├── quick-start.md
│   ├── user-guide.md
│   └── faq.md
├── developer/
│   ├── installation.md
│   ├── architecture.md
│   ├── api-reference.md
│   └── operations.md
└── images/
    ├── interface.svg
    ├── architecture.svg
    └── deployment-flow.svg
```

## Просмотр

- **VS Code:** откройте папку, `Ctrl+Shift+V` на нужном `.md`-файле.
- **GitHub/GitLab:** загрузите папку в репозиторий — ссылки и изображения работают сразу.
- **MkDocs:** `pip install mkdocs-material`, затем `mkdocs serve` и откройте http://127.0.0.1:8000
