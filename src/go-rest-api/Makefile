all: swag build

swag:
	swag init -g cmd/server/main.go --parseInternal --parseDependency

build:
	go build -o ./server cmd/server/main.go

.PHONY: all swag build
