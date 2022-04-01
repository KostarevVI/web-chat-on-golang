.PHONY: build
build:
	go build -v ./cmd/chatserver

.DEFAULT_GOAL := build