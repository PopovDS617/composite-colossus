FROM golang:1.21.6.3-alpine AS builder

COPY . /app/
WORKDIR /app/

RUN go mod download
RUN go build -o ./bin/server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/bin/server .
COPY --from=builder /app/local.env .

CMD ["./server", "-config-path=local.env"]