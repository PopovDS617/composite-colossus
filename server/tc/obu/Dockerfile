FROM golang:1.21.6-alpine

WORKDIR /app

COPY . .

RUN go mod tidy


RUN go build -o main .

ENV WS_RECEIVER_ADDRESS="ws://data_receiver:30000/ws"

CMD ["./main"]