run:
	go run ./cmd/api

build:
	go build ./cmd/api

.DEFAULT_GOAL: run