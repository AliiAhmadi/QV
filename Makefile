# Ali Ahmadi (C) 2023

.DEFAULT_GOAL := help

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n "s/^##//p" ${MAKEFILE_LIST} | column -t -s ':' | sed -e "s/^/ /"

## run: run the application
.PHONY: run
run:
	@go run .

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: fmt vet test

## fmt: formatting codes
.PHONY: fmt
fmt:
	@echo "formatting..."
	@go fmt ./...

## vet: vetting codes
.PHONY: vet
vet:
	@echo "vetting..."
	@go vet ./...

## test: run all tests
.PHONY: test
test:
	@echo "running tests..."
	@go test -race -vet=off ./...

linker_flag = "-s"

## build: build the application
.PHONY: build
build:
	@echo "building QV..."
	GOOS=windows GOARCH=amd64 go build -a -ldflags=${linker_flag} -o=./bin/windows_amd64/QV .
	GOOS=linux GOARCH=amd64 go build -a -ldflags=${linker_flag} -o=./bin/linux_amd64/QV .
	GOOS=linux GOARCH=mips64 go build -a -ldflags=${linker_flag} -o=./bin/linux_mips64/QV .
	GOOS=linux GOARCH=arm64 go build -a -ldflags=${linker_flag} -o=./bin/linux_arm64/QV .

## clear: clear cache for tests
.PHONY: clear
clear:
	@go clean -testcache