FROM golang:1.15

WORKDIR /var/www/src

RUN apt update
RUN go mod init wing
RUN go mod tidy
RUN go get -u github.com/cosmtrek/air
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate.linux-amd64 /usr/bin/migrate
ENV MYSQL_URL=mysql://root:root@tcp(mysql:3306)/wing
