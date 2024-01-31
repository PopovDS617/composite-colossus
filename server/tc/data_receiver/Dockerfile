FROM golang:1.21.6-alpine as builder

WORKDIR /app

COPY . .

RUN apk update && \
apk add git gcc libc-dev

RUN go mod tidy

ENV CGO_ENABLED 1
ENV GOFLAGS -mod=readonly
ENV GOOS=linux
ENV GOARCH=amd64

RUN export GO111MODULE=on

RUN go build -tags musl -o main .

CMD ["./main"]