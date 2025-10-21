GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)

all: benchmark benchmark_client

benchmark:
	go build -o $(GOBIN)/benchmark $(ROOT)/cmd/benchmark
	@echo "Done building benchmark."

benchmark_client:
	go build -o $(GOBIN)/benchmark_client $(ROOT)/cmd/client
	@echo "Done building benchmark_client."

clean:
	rm -rf $(ROOT)/build