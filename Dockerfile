FROM golang:1.21.1

WORKDIR /app

ENV CONFIG=docker

COPY .. /app


RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go mod download


ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command=./main