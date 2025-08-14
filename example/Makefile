GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)
all:blacklist coupon election lottery oracle upgrade clean emptynode

blacklist:blacklist.bin
coupon:coupon.bin
election:election.bin
lottery:lottery.bin
oracle:oracle.bin
upgrade:upgrade.bin
emptynode:emptynode.bin
%.bin:
	@if [ -d "$(ROOT)/$*/cmd" ]; then \
		echo "Building $* from cmd/..."; \
		go build -o $(GOBIN)/$* $(ROOT)/$*/cmd; \
	else \
		echo "Building $* from project root..."; \
		go build -o $(GOBIN)/$* $(ROOT)/$*; \
	fi
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(*)\" to launch $(*)."

clean:
	rm -rf $(ROOT)/build
