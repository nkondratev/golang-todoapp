#!/usr/bin/env bash
for i in {1..100}; do
    curl -o /dev/null -s -w "%{time_total}\n" \
    http://127.0.0.1:5050/api/v1/tasks
done
