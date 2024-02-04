FROM golang:1.21.6-alpine

WORKDIR /app

COPY . .

RUN go mod tidy


RUN go build -o main .

ENV PORT="10000"
ENV AGGREGATOR_ADDRESS="http://data_aggregator:9000"

CMD ["./main"]