#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

migrate \
  -path "$PROJECT_ROOT/db/migrations" \
  -database "postgres://postgres:password@localhost:5433/urlshortener?sslmode=disable" \
  up