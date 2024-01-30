FROM golang:1.21.6-alpine

WORKDIR /app

COPY . .

RUN go mod tidy


RUN go build -o main .

EXPOSE 9000

CMD ["./main"]