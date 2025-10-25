.DEFAULT_GOAL := build

.PHONY: fmt vet build

fmt:
	go fmt ./...

vet:fmt
	go vet ./...

build:vet
	go build ./cmd/app

clean:
	rm -f build

tidy:
	go mod tidy

test:
	go test ./internal/task/task_handle_test -v