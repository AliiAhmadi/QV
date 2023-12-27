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