GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)
.PHONY:tools simapp example clean
all:tools simapp example

tools:tools.bin
simapp:
	go build -o $(GOBIN)/simapp $(ROOT)/simapp
	go build -o $(GOBIN)/simapp_client $(ROOT)/simapp/cmd/client
	@echo "Done building."
	@echo "Run \"$(GOBIN)/simapp\" to launch simapp."
	@echo "Run \"$(GOBIN)/simapp_client\" to launch simapp_client."

example:
	make -f example/Makefile GOBIN=$(ROOT)/example/build/bin ROOT=$(ROOT)/example

%.bin:
	go build -o $(GOBIN)/$* $(ROOT)/$*/cmd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(*)\" to launch $(*)."
clean:
	rm -rf $(ROOT)/build
