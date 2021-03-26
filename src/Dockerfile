FROM golang:latest

WORKDIR /usr/src/app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

CMD ["go", "run", "main.go"]
