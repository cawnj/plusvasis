# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./server ./cmd/server/main.go

# Build the server image

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY --from=build /app/server ./server

EXPOSE 8080
CMD [ "./server" ]

