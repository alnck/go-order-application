FROM golang:1.17

RUN mkdir /customerApp
WORKDIR /customerApp

COPY /customer-service .

RUN go mod download

RUN go build ./src/main.go

EXPOSE 5001

CMD ["./main"]