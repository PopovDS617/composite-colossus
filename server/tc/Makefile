obu:
	sleep 2
	@go build -o bin/obu ./obu/main.go
	@./bin/obu

rcvr:
	@go build -o bin/data_receiver ./data_receiver/main.go
	@./bin/data_receiver

dcalc:
	@go build -o bin/distance_calculator ./distance_calculator/main.go
	@./bin/distance_calculator

daggr:
	@go build -o bin/data_aggregator ./data_aggregator/main.go
	@./bin/data_aggregator
	
runall: dcalc rcvr obu daggr


proto:
	PATH="${PATH}:${HOME}/go/bin" protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./distance_calculator/pb/ptypes.proto
	PATH="${PATH}:${HOME}/go/bin" protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./data_aggregator/pb/ptypes.proto


.PHONY: obu