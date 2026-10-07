#!/usr/bin/env bash
ab -n 1000 -c 100 http://127.0.0.1:5050/api/v1/tasks
