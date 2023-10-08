FROM golang:latest

ARG WOW_ROOT_DIR
ARG WOW_SERVER_CONFIG

ENV WOW_SERVER_CONFIG=${WOW_SERVER_CONFIG}
ENV WOW_ROOT_DIR=/app

WORKDIR /app

COPY . .

RUN go build -o server ./cmd/server/main.go

CMD ["./server"]