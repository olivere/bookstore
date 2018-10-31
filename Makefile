default: build

.PHONY: build
build:
	go build ./cmd/bookctl

.PHONY: web
web:
	go build ./cmd/bookweb
