FROM golang:1.17

RUN mkdir /orderApp
WORKDIR /orderApp

COPY /order-service .

RUN go mod download

RUN go build ./src/main.go

EXPOSE 5002

CMD ["./main"]