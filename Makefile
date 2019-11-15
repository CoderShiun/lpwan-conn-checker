.PHONY: build clean
VERSION := $(shell git describe --tags |sed -e "s/^v//")

build: #scripts
	mkdir -p build
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/conn-checker main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build

#scripts:
#	@sh $(shell pwd)/scripts/sync_env_variables
