SHELL:=/bin/bash -O extglob
BINARY=ms-api
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"


up:
	docker-compose up --build

down:
	docker-compose down --remove-orphans

restart: down up