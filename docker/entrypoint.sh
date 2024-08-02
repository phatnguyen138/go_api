#!/bin/sh
set -e

/usr/local/bin/migrate -path internal/db/migration -database "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable" -verbose up

exec "$@"