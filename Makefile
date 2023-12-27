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