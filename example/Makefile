GOBIN = $(shell pwd)/build/bin
ROOT=$(shell pwd)
all:blacklist coupon election lottery oracle upgrade

blacklist:blacklist.bin
coupon:coupon.bin
election:election.bin
lottery:lottery.bin
oracle:oracle.bin
upgrade:upgrade.bin

%.bin:
	go build -o $(GOBIN)/$* $(ROOT)/$*/cmd
	@echo "Done building."
	@echo "Run \"$(GOBIN)/$(*)\" to launch $(*)."