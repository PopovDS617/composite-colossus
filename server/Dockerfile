FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .



ENV HTTP_LISTEN_ADDRESS=:5000

ENV JWT_SECRET=secretsecretsecretsecretsecretsecretsecretsecret

ENV MONGO_DB_NAME=hotel-reservation
ENV MONGO_DB_URI=mongodb://root:password@mongodb/?authSource=admin

ENV MONGO_DB_NAME_TEST=hotel-reservation-test
ENV MONGO_DB_URI_TEST=mongodb://root:password@mongodb/?authSource=admin

EXPOSE 5000

CMD ["./main"]