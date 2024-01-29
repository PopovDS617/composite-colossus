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



.PHONY: obu