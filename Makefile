.DEFAULT_GOAL: build

.PHONY: fmt vet build

.fmt:
	fmt ./...

.vet:fmt
	vet ./...

.build:vet
	go build