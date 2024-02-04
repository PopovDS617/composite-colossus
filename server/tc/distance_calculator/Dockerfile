FROM golang:1.21.6-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN apk update && \
apk add git gcc libc-dev


ENV CGO_ENABLED 1
ENV GOFLAGS -mod=readonly
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on

RUN go build -tags musl -o main .

ENV AGGREGATOR_ADDRESS="http://data_aggregator:9000/aggregator" 
ENV BOOTSTRAP_SERVERS="kafka:29092"

CMD ["./main"]