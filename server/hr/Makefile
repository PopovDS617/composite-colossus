build:
	go build -o bin/api

run: build
	./bin/api

test:
	go test -v ./... -count=1

seed:
	@go run scripts/seed.go

docker:
	@docker build -t app .
	@docker run -p 5000:5000 app