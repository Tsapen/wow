FROM golang:latest

ARG WOW_ROOT_DIR
ARG WOW_CLIENT_CONFIG

ENV WOW_CLIENT_CONFIG=${WOW_CLIENT_CONFIG}
ENV WOW_ROOT_DIR=/app

WORKDIR /app

COPY . .

RUN go build -o client ./cmd/client/main.go

CMD ["./client"]
