#!/bin/bash

set -e

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo "Ошибка: файл .env не найден"
    exit 1
fi

BACKUP_DIR="${PROJECT_ROOT}/out/backups"
DATE=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="${BACKUP_DIR}/todoapp_${DATE}.dump"

mkdir -p "$BACKUP_DIR"

echo "Создание резервной копии базы данных..."
echo "База данных: $POSTGRES_DB"

docker compose exec -T todoapp-postgres \
    pg_dump \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB" \
    -Fc > "$BACKUP_FILE"

echo "Резервная копия успешно создана:"
echo "$BACKUP_FILE"
