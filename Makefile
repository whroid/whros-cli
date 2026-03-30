.PHONY: build install clean test lint

NAME=whros
VERSION=0.1.0
BUILD_DIR=./bin

build:
	go build -o $(BUILD_DIR)/$(NAME) main.go

install: build
	cp $(BUILD_DIR)/$(NAME) /usr/local/bin/

clean:
	rm -rf $(BUILD_DIR)

test:
	go test ./...

lint:
	golangci-lint run

dev:
	go run main.go