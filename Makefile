all: build run

build:
	docker build -t pong .

run:
	docker run --rm -it pong

.PHONY: build run
