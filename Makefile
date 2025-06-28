APP      ?= pack-calculator
PORT     ?= :8081
IMAGE    := $(APP):latest
BIN_DIR  := bin
BIN_FILE := $(BIN_DIR)/$(APP)

.DEFAULT_GOAL := help

.PHONY: help start compose build test clean docker-build

start:
	@echo "⇢ starting $(APP) on $(PORT)"
	go run ./cmd/server

compose:
	docker compose up --build

build:
	@echo "→ building $(BIN_FILE)…"
	@mkdir -p $(BIN_DIR)
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(BIN_FILE) ./cmd/server

test:
	go test -v ./...

docker-build:
	docker build -t $(IMAGE) .

docker-run: docker-build
	docker run --rm -p 8081:8081 \
		-e PACKCALC_PORT=$(PORT) \
		$(IMAGE)

clean:
	rm -rf $(BIN_DIR)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	  awk 'BEGIN {FS = ":.*?## "}; {printf " \033[36m%-14s\033[0m %s\n", $$1, $$2}'