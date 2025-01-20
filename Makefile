GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)
.PHONY:tools simapp example
all:tools simapp example

tools:tools.bin
simapp:
	go build -o $(GOBIN)/simapp $(ROOT)/simapp
	@echo "Done building."
	@echo "Run \"$(GOBIN)/simapp\" to launch simapp."
example:
	make -f example/Makefile GOBIN=$(ROOT)/example/build/bin ROOT=$(ROOT)/example

%.bin:
	go build -o $(GOBIN)/$* $(ROOT)/$*/cmd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(*)\" to launch $(*)."