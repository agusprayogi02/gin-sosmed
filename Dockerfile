FROM mysql:8.0 # Official MySQL image

# -- Optional (persistent data)
# VOLUME /var/lib/mysql

RUN apt-get update && apt-get install -y --no-install-recommends \
    mariadb-server-client

# -- Optional (set root password)
ENV MYSQL_ROOT_PASSWORD=root1234
RUN echo "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root1234' WITH GRANT OPTION;" > /tmp/mysql_setup.sql && mysql -u root < /tmp/mysql_setup.sql && rm /tmp/mysql_setup.sql

# -- Your Golang application build stages below

FROM golang:1.21-alpine

RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod go.sum ./
COPY .env.example .env ./

RUN go mod download

COPY . .

RUN go build -o main

CMD ["./main"]