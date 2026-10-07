#!/bin/bash

PROJECT_DIR="${1:-.}"

cd "$PROJECT_DIR" || exit 1

find . \
    -type f \
    -not -path './.git/*' \
    -not -path './out/logs/*' \
    -not -path './out/pgdata/*' \
    -print0 |
    sort -z |
    while IFS= read -r -d '' file; do
        sha256sum "$file"
    done |
    sha256sum
