FROM golang:1.22-alpine

RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod go.sum ./
COPY .env.example .env ./

RUN go mod tidy

COPY . .

RUN go build -o main

CMD ["dockerize", "-wait", "tcp://locahost:3306", "-timeout", "20s", "./main"]