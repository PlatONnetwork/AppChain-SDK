GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)
all:benchmark client

benchmark:benchmark.bin
client:client.bin
%.bin:
	go build -o $(GOBIN)/$* $(ROOT)/cmd/$*
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(*)\" to launch $(*)."

clean:
	rm -rf $(ROOT)/build