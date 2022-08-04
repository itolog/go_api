run:
	go run ./cmd/api

build:
	go build ./cmd/api

lint:
	golangci-lint run ./... -v

lint_fix:
	golangci-lint run ./... -v --fix

.DEFAULT_GOAL: run