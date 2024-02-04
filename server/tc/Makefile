obu:
	sleep 2
	@go build -o bin/obu ./obu/main.go
	@WS_RECEIVER_ADDRESS=ws://localhost:30000/ws ./bin/obu

rcvr:
	@go build -o bin/data_receiver ./data_receiver/main.go
	@BOOTSTRAP_SERVERS=localhost:9092 PORT=30000 ./bin/data_receiver

dcalc:
	@go build -o bin/distance_calculator ./distance_calculator/main.go
	@AGGREGATOR_ADDRESS=http://localhost:9000/aggregator BOOTSTRAP_SERVERS=localhost:9092 ./bin/distance_calculator

daggr:
	@go build -o bin/data_aggregator ./data_aggregator/main.go
	@HTTP_PORT=9000 GRPC_PORT=9001 ./bin/data_aggregator
	
gateway:
	@go build -o bin/gateway ./gateway/main.go
	@PORT=10000 AGGREGATOR_ADDRESS=http://localhost:9000 ./bin/gateway

prome:
	docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus

runall: gateway dcalc rcvr obu daggr


proto:
	PATH="${PATH}:${HOME}/go/bin" protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./distance_calculator/pb/ptypes.proto
	PATH="${PATH}:${HOME}/go/bin" protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./data_aggregator/pb/ptypes.proto
	PATH="${PATH}:${HOME}/go/bin" protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./gateway/pb/ptypes.proto


.PHONY: obu gateway