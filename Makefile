run:
	go run ./cmd/api

build:
	go build ./cmd/api

lint:
	golangci-lint run ./... -v	

.DEFAULT_GOAL: run