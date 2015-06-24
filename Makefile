all: test build run

build:
	docker build -t pong .

run:
	docker run --rm -it pong

test:
	go test ./...

.PHONY: build run test
