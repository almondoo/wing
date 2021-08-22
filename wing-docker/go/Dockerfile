FROM golang:1.17

WORKDIR /var/www/src

RUN apt update
RUN go mod init wing
RUN go mod tidy
RUN go get -u github.com/cosmtrek/air
ENV MYSQL_URL=mysql://root:root@tcp(mysql:3306)/wing
