# syntax=docker/dockerfile:1

FROM golang:1.20.4-alpine3.17

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod .
COPY ./cmd ./cmd/
COPY ./internal ./internal/

RUN go mod download

RUN go build -o server ./cmd/server/

EXPOSE 8080

CMD ["./server"]
