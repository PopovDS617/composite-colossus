FROM golang:1.21.6-alpine

WORKDIR /app

COPY . .

RUN go mod tidy


RUN go build -o main .

EXPOSE 9000

ENV HTTP_PORT="9000" 
ENV GRPC_PORT="9001"

CMD ["./main"]